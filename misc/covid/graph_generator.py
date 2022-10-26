"""this python script is a cmdline program that generates graphs from data.
This is hastily written and is just meant to produce graphs from data located
in the data_exp folder"""

from __future__ import annotations
from cmath import exp
from enum import Enum
import itertools

from typing import Dict, List, Optional, Set, Tuple
import collections

from matplotlib.axes import Axes


class Line(Enum):
    EXP_TYPE = 0
    TRIAL = 1
    THREADS = 2
    INTERRUPT = 3
    REAL = 4
    USER = 5
    SYS = 6
    ZIPCODE = 7
    MONTH = 8
    YEAR =  9
    UNKNOWN = 10

class ExpType(Enum):
    SEQUENTIAL = 0
    STATIC = 1
    MAP = 2
    UNKNOWN = 3

class Experiment:

    def __init__(self, trials: List[Trial]):
        
        self._map_experiments: List[Trial] = []
        self._seq_experiments: List[Trial] = []
        self._static_experiments: List[Trial] = []

        for trial in trials:
            if trial.exp_type == ExpType.MAP:
                self._map_experiments.append(trial)
            elif trial.exp_type == ExpType.SEQUENTIAL:
                self._seq_experiments.append(trial)
            elif trial.exp_type == ExpType.STATIC:
                self._static_experiments.append(trial)

    def get_avg_seq_time(self, seq_trials: List[Trial]) -> float:
        """
        returns the avg wallclock time (secs) of the sequential solve,
        given a list of trials conducted for the sequential solve.
        """
        times = [trial.real_sec for trial in seq_trials]
        return sum(times)/len(times)

    def get_x_s_serr_data(self, trials: List[Trial], seq_trials: List[Trial]) -> collections.OrderedDict[int,Tuple[float,float,float]]:

        avg_seq_time = self.get_avg_seq_time(seq_trials) # avg sequential time
        x_y_yerr_data_dict = self.get_x_y_yerr_data(trials)   # parallel times
        
        # y, y_errormin, y_errormax = zip(*x_y_yerr_data_dict.values())
        xy_dict = self.get_x_y_data(trials)

        x_s_serr_data: collections.OrderedDict[int,Tuple[float,float,float]] = dict()
        for key in sorted(xy_dict.keys()):
            x_s_serr_data[key] = tuple()

        for key in xy_dict.keys():
            avg_time = sum(xy_dict[key])/len(xy_dict[key])
            avg_s = self.speed_up(avg_seq_time,avg_time)
            # max time -> min speed up
            max_time = max(xy_dict[key])
            min_s = self.speed_up(avg_seq_time,max_time)
            s_bot_err = avg_s - min_s
            # min time -> max speed up
            min_time = min(xy_dict[key])
            max_s = self.speed_up(avg_seq_time,min_time)
            s_top_err = max_s - avg_s
            
            x_s_serr_data[key] = (avg_s,s_bot_err,s_top_err)

        return x_s_serr_data

    def speed_up(self, seq_time: float,parallel_time: float) -> float:
        return seq_time/parallel_time

    def get_x_y_data(self, trials: List[Trial]) -> Dict[int,List[float]]:
        """
        accepts a list of trials and returns a dictionary mapping number of threads (x value)
        to a list wallclock times (y values) for every trial conducted at that thread count
        """

        unique_thread_cnts: Set[int] = set()

        for trial in trials:
            unique_thread_cnts.add(trial.threads)

        threads_cnts = list(unique_thread_cnts)

        xy_dict: Dict[int,List[float]] = dict()

        for thread_cnt in threads_cnts:
            xy_dict[thread_cnt] = []

        for trial in trials:
            xy_dict[trial.threads].append(trial.real_sec)

        return xy_dict

    def get_x_y_yerr_data(self, trials: List[Trial]) -> collections.OrderedDict[int,Tuple[float,float,float]]:

        xy_dict = self.get_x_y_data(trials)

        # maps x value to (y_avg, y_min, y_max)
        x_y_yerr_data_dict: collections.OrderedDict[int,Tuple[float,float,float]] = dict()
        for key in sorted(xy_dict.keys()):
            x_y_yerr_data_dict[key] = tuple()

        for key in xy_dict.keys():
            avg_time = sum(xy_dict[key])/len(xy_dict[key])
            max_time = max(xy_dict[key])
            y_top_err = max_time - avg_time
            min_time = min(xy_dict[key])
            y_bot_err = avg_time - min_time
            x_y_yerr_data_dict[key] = (avg_time,y_bot_err,y_top_err)

        return x_y_yerr_data_dict

    def graph_speedup_trials_all(self, abs_write_dir_loc: str):
        all_trials = list(itertools.chain(self._map_experiments, self._seq_experiments, self._static_experiments))
        self.graph_speedup_trials(all_trials,abs_write_dir_loc)

    def graph_speedup_trials(self, trials: List[Trial], abs_write_dir_loc: str):
        """
        graphs speedup of a set of trials compared to a set of baseline
        trials from sequential solving
        """

        import numpy as np
        import matplotlib.pyplot as plt

        # add line graphs
        seq_experiments = [trial for trial in trials if trial.exp_type==ExpType.SEQUENTIAL]
        static_experiments = [trial for trial in trials if trial.exp_type==ExpType.STATIC]
        map_experiments = [trial for trial in trials if trial.exp_type==ExpType.MAP]

        fig, ax = plt.subplots()
        fig.set_size_inches(10, 5)
        
        self.add_speedup_line_graph(ax, static_experiments, seq_experiments)
        self.add_speedup_line_graph(ax, map_experiments, seq_experiments)

        ax.set_xlabel('number of threads')
        ax.set_ylabel('speed up (relative to average sequential time)')
        ax.set_title(f'Speed up relative to average sequential time')
        ax.set_xscale('log')
        ax.legend(loc="lower right")
        

        fig_path = os.path.join(abs_write_dir_loc,'speedup.pdf')
        fig.savefig(fig_path)
        print(f"generated and saved figure to: {fig_path}")
        # plt.show()
        # fig.savefig('fig2.pdf')
        plt.show()

    def add_speedup_line_graph(self, ax: Axes, trials: List[Trial], seq_trials: List[Trial]):
        exp_type = trials[0].exp_type
        zipcode = trials[0].zipcode
        month = trials[0].month
        year = trials[0].year
        num_trials = max([trial.trial for trial in trials])
        x_s_serr_data_dict = self.get_x_s_serr_data(trials, seq_trials)
        x = x_s_serr_data_dict.keys()    
        s, s_errormin, s_errormax = zip(*x_s_serr_data_dict.values())
        ax.errorbar(x, s,
            # xerr=xerr,
            yerr=[s_errormin,s_errormax],
            fmt='-o',
            label = f"method={exp_type}, zipcode={zipcode}, month={month}, year={year}, num_trials={num_trials}")
        
    def add_process_line_graph(self, ax: Axes, trials: List[Trial]):
        exp_type = trials[0].exp_type
        zipcode = trials[0].zipcode
        month = trials[0].month
        year = trials[0].year
        num_trials = max([trial.trial for trial in trials])
        x_y_yerr_data_dict = self.get_x_y_yerr_data(trials)
        x = x_y_yerr_data_dict.keys()
        y, y_errormin, y_errormax = zip(*x_y_yerr_data_dict.values())
        ax.errorbar(x, y,
                    # xerr=xerr,
                    yerr=[y_errormin,y_errormax],
                    fmt='-o',
                    label = f"method={exp_type}, zipcode={zipcode}, month={month}, year={year}, num_trials={num_trials}")

    def graph_process_trials_all(self, abs_write_dir_loc: str):
        """
        graphs the total process time of solving the covid data wrangling problem
        as a function of the number of the number of threads used for all methods
        combined that are part of this experiment.
        """
        all_trials = list(itertools.chain(self._map_experiments, self._seq_experiments, self._static_experiments))
        self.graph_process_trials(all_trials,abs_write_dir_loc)

    def graph_process_trials(self, trials: List[Trial], abs_write_dir_loc: str):
        """
        accepts a list of Trial objects and graphs the total process time of
        solving the covid data wrangling problem as a function of the number of
        the number of threads used.
        """
        import numpy as np
        import matplotlib.pyplot as plt

        fig, ax = plt.subplots()
        fig.set_size_inches(10, 5)

        # add line graphs
        seq_experiments = [trial for trial in trials if trial.exp_type==ExpType.SEQUENTIAL]
        static_experiments = [trial for trial in trials if trial.exp_type==ExpType.STATIC]
        map_experiments = [trial for trial in trials if trial.exp_type==ExpType.MAP]

        self.add_process_line_graph(ax, seq_experiments)
        self.add_process_line_graph(ax, static_experiments)
        self.add_process_line_graph(ax, map_experiments)

        ax.set_xlabel('number of threads')
        ax.set_ylabel('total process time [secs]')
        ax.set_title('total process time as fcn of number of threads')
        ax.set_xscale('log')
        ax.legend(loc="upper right")

    
        fig_path = os.path.join(abs_write_dir_loc,'process_time.pdf')
        fig.savefig(fig_path)
        print(f"generated and saved figure to: {fig_path}")


        plt.show()


class Trial:

    def __init__(self, abs_filepath: str):

        self.exp_type: ExpType = ExpType.UNKNOWN
        self.real_sec: Optional[float] = None
        self.trial: int = None
        self.threads: int = None
        self.zipcode: str = None
        self.month: int = None
        self.year: int = None
        
        with open(abs_filepath) as f:
            lines = f.readlines()
        
        lines = [line.strip() for line in lines]

        for line in lines:

            line_type = self._characterize_line(line)

            # if line_type in [Line.EXP_TYPE, Line.TRIAL, Line.THREADS, Line.REAL]:
            try:
                self._analyze_line(line,line_type)
            except Exception as e:
                pass # do not errors to console; just silence them.
                # print(e)
            
            if line_type == Line.INTERRUPT: # signal to disregard this trial
                raise Exception("interrupt signal encountered")

    def _characterize_line (self, line: str) -> Line:
        """characterizes and returns the type of line"""
        if line == 'signal: interrupt':
            return Line.INTERRUPT
        elif line[0:4] == "real":
            return Line.REAL
        elif line[0:4] == "user":
            return Line.USER
        elif line[0:3] == "sys":
            return Line.SYS
        elif line[0:5] == "TRIAL":
            return Line.TRIAL
        elif line[0:7] == "THREADS":
            return Line.THREADS
        elif line[0:8] == "EXP_TYPE":
            return Line.EXP_TYPE
        elif line[0:7] == "ZIPCODE":
            return Line.ZIPCODE
        elif line[0:5] == "MONTH":
            return Line.MONTH
        elif line[0:4] == "YEAR":
            return Line.YEAR
        else:
            return Line.UNKNOWN

    def _analyze_line(self,line: str, line_type: Line):
        """
        processes a line according to its line type.
        modifies internal variables
        """

        # extract time
        if line_type == Line.REAL:
            try:
                num_str = line[7:12]
                self.real_sec = float(num_str)
            except:
                print("failed to interpret time as float: '",num_str, "' from line: ",line)

        # extract experiment type
        elif line_type == Line.EXP_TYPE:
            try:
                exp_type_str = line[9:]
                if exp_type_str == "STATIC":
                    self.exp_type = ExpType.STATIC
                elif exp_type_str == "MAP":
                    self.exp_type = ExpType.MAP
                elif exp_type_str == "SEQUENTIAL":
                    self.exp_type = ExpType.SEQUENTIAL
                else:
                    print("failed to interpret experiment type: ",exp_type_str, "from line: ",line)
            except:
                print("failed to interpret time as float: ",num_str, "from line: ",line)
        
        # extract trial number
        elif line_type == Line.TRIAL:
            try:
                trial_str = line[6:]
                self.trial = int(trial_str)
            except:
                print("failed to interpret trial number: ",trial_str, "from line: ",line)

        # extract number of threads
        elif line_type == Line.THREADS:
            try:
                threads_str = line[8:]
                self.threads = int(threads_str)
            except:
                print("failed to interpret number of threads: ",threads_str, "from line: ",line)

        # extract zipcode
        elif line_type == Line.ZIPCODE:
            try:
                zipcode_str = line[8:]
                self.zipcode = zipcode_str
            except:
                print("failed to interpret zipcode: ",zipcode_str, "from line: ",line)

        # extract month
        elif line_type == Line.MONTH:
            try:
                month_str = line[6:]
                self.month = int(month_str)
            except:
                print("failed to interpret month: ",month_str, "from line: ",line)

        # extract year
        elif line_type == Line.YEAR:
            try:
                year_str = line[5:]
                self.year = int(year_str)
            except:
                print("failed to interpret year: ",year_str, "from line: ",line)

        else:
            raise ValueError("line_type not accepted for analysis: ",line_type)

    def __repr__(self) -> str:
        return f"Experiment(type={self.exp_type}, trial={self.trial}, threads={self.threads}, time={self.real_sec} secs, " + \
            f"zip={self.zipcode}, month={self.month}, year={self.year})"


if __name__ == "__main__":

    import os

    path_to_this_dir = os.path.dirname(os.path.abspath(__file__))
    path_to_expmnts_dir = os.path.join(path_to_this_dir,"data_exp")
    files = os.listdir(path_to_expmnts_dir)
    most_recent_exp_name = max(files)
    path_to_exp_data_dir = os.path.join(path_to_expmnts_dir,most_recent_exp_name)

    print("analyzing data from: ",path_to_exp_data_dir)

    data_files = os.listdir(path_to_exp_data_dir)
    print(f"found {len(data_files)} files (data points)")

    abs_paths_to_data_files = [os.path.join(path_to_exp_data_dir,filename) for filename in data_files]

    # create Trial objects from each file
    experiment_trials: List[Trial] = []

    for filepath in abs_paths_to_data_files:
        try:
            exp_trial = Trial(filepath)
            experiment_trials.append(exp_trial)
        except Exception as e:
            pass # fail silently
            # print("FAILED to extract experiment info.")
            # print(e)

    print(f"extracted info from {len(experiment_trials)} experiments")
    experiment = Experiment(experiment_trials)

    # print()
    # print(experiment_trials[0])
    # print()
    # print(experiment._seq_experiments)
    # print()
    # print(experiment._static_experiments)
    # print()
    # print(experiment._map_experiments)
    # print()
    # print(experiment.get_x_y_yerr_data(experiment._static_experiments))
    # print()
    # print(experiment.get_x_s_serr_data(experiment._static_experiments, experiment._seq_experiments))

    experiment.graph_speedup_trials_all(path_to_exp_data_dir)
    experiment.graph_process_trials_all(path_to_exp_data_dir)

