"""this python script is a cmdline program that generates graphs from data.
This is hastily written and is just meant to produce graphs from data located
in the data_exp folder"""

from __future__ import annotations
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
    MIN_DELAY = 7
    MAX_DELAY = 8
    INCREMENT_AMOUNT =  9
    UNKNOWN = 10

class ExpType(Enum):
    TAS = 0
    TTAS = 1
    EB = 2
    EBF = 3
    SYNC = 4
    UNKNOWN = 3

class Experiment:

    def __init__(self, trials: List[Trial], announce_findings: bool = False):
        
        self._tas_experiments: List[Trial] = []
        self._ttas_experiments: List[Trial] = []
        self._eb_experiments: List[Trial] = []
        self._ebf_experiments: List[Trial] = []
        self._sync_experiments: List[Trial] = []

        for trial in trials:
            if trial.exp_type == ExpType.TAS:
                self._tas_experiments.append(trial)
            elif trial.exp_type == ExpType.TTAS:
                self._ttas_experiments.append(trial)
            elif trial.exp_type == ExpType.EB:
                self._eb_experiments.append(trial)
            elif trial.exp_type == ExpType.EBF:
                self._ebf_experiments.append(trial)
            elif trial.exp_type == ExpType.SYNC:
                self._sync_experiments.append(trial)

        if announce_findings:
            print("TAS  experiments: ",len(self._tas_experiments))
            print("TTAS experiments: ",len(self._ttas_experiments))
            print("EB   experiments: ",len(self._eb_experiments))
            print("EBF  experiments: ",len(self._ebf_experiments))
            print("SYNC experiments: ",len(self._sync_experiments))

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
        
    def add_process_line_graph(self, ax: Axes, trials: List[Trial]):
        exp_type = trials[0].exp_type
        min_delay = trials[0].min_delay
        max_delay = trials[0].max_delay
        inc_amount_str = "{:.0e}".format(trials[0].increment_amnt)
        num_trials = max([trial.trial for trial in trials])
        x_y_yerr_data_dict = self.get_x_y_yerr_data(trials)
        x = x_y_yerr_data_dict.keys()
        y, y_errormin, y_errormax = zip(*x_y_yerr_data_dict.values())

        if exp_type in [ExpType.EB,ExpType.EBF]: 
            label = f"locktype={exp_type}, min_delay={min_delay}, max_delay={max_delay}, inc_amnt={inc_amount_str}, num_trials={num_trials}"
        else:
            label = f"locktype={exp_type}, inc_amnt={inc_amount_str}, num_trials={num_trials}"

        ax.errorbar(x, y,
                    # xerr=xerr,
                    yerr=[y_errormin,y_errormax],
                    fmt='-o',
                    label = label)

    def graph_process_trials_all(self, abs_write_dir_loc: str):
        """
        graphs the total process time of running experiments on lock implementations
        as a function of the number of the number of threads used with that particular
        lock implementation. I.e., there is one line graph for each lock implementation.
        """
        all_trials = list(itertools.chain(
            self._tas_experiments,self._ttas_experiments,
            self._eb_experiments,self._ebf_experiments,
            self._sync_experiments
            ))
        self.graph_process_trials(all_trials,abs_write_dir_loc)

    def graph_process_trials(self, trials: List[Trial], abs_write_dir_loc: str):
        """
        accepts a list of Trial objects and graphs the total process time of
        running experiments on lock implementations as a function of the number of
        the number of threads used.
        """
        import numpy as np
        import matplotlib.pyplot as plt

        fig, ax = plt.subplots()
        fig.set_size_inches(10, 12)

        # add line graphs
        tas_experiments = [trial for trial in trials if trial.exp_type==ExpType.TAS]
        ttas_experiments = [trial for trial in trials if trial.exp_type==ExpType.TTAS]
        eb_experiments = [trial for trial in trials if trial.exp_type==ExpType.EB]
        ebf_experiments = [trial for trial in trials if trial.exp_type==ExpType.EBF]
        sync_experiments = [trial for trial in trials if trial.exp_type==ExpType.SYNC]

        self.add_process_line_graph(ax, tas_experiments)
        self.add_process_line_graph(ax, ttas_experiments)
        self.add_process_line_graph(ax, eb_experiments)
        self.add_process_line_graph(ax, ebf_experiments)
        self.add_process_line_graph(ax, sync_experiments)

        ax.set_xlabel('number of threads')
        ax.set_ylabel('total process time [secs]')
        ax.set_title('total process time as fcn of number of threads')
        ax.legend(loc="upper left")

    
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
        self.min_delay: str = None # ms
        self.max_delay: int = None # ms
        self.increment_amnt: int = None
        
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
        elif line[0:9] == "MIN_DELAY":
            return Line.MIN_DELAY
        elif line[0:9] == "MAX_DELAY":
            return Line.MAX_DELAY
        elif line[0:16] == "INCREMENT_AMOUNT":
            return Line.INCREMENT_AMOUNT
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
                if exp_type_str == "TAS":
                    self.exp_type = ExpType.TAS
                elif exp_type_str == "TTAS":
                    self.exp_type = ExpType.TTAS
                elif exp_type_str == "EB" and exp_type_str != "EBF":
                    self.exp_type = ExpType.EB
                elif exp_type_str == "EBF":
                    self.exp_type = ExpType.EBF
                elif exp_type_str == "TTAS":
                    self.exp_type = ExpType.TTAS
                elif exp_type_str == "SYNC":
                    self.exp_type = ExpType.SYNC
                else:
                    print("failed to interpret experiment type: ",exp_type_str, "from line: ",line)
                    self.exp_type = ExpType.UNKNOWN
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

        # extract min_delay
        elif line_type == Line.MIN_DELAY:
            try:
                min_delay_str = line[10:]
                self.min_delay = int(min_delay_str)
            except:
                print("failed to interpret min_delay: ",min_delay_str, "from line: ",line)

        # extract max_delay
        elif line_type == Line.MAX_DELAY:
            try:
                max_delay_str = line[10:]
                self.max_delay = int(max_delay_str)
            except:
                print("failed to interpret max_delay: ",max_delay_str, "from line: ",line)

        # extract year
        elif line_type == Line.INCREMENT_AMOUNT:
            try:
                inc_amnt_str = line[17:]
                self.increment_amnt = int(inc_amnt_str)
            except:
                print("failed to interpret increment amount: ",inc_amnt_str, "from line: ",line)

        else:
            raise ValueError("line_type not accepted for analysis: ",line_type)

    def __repr__(self) -> str:
        return f"Experiment(type={self.exp_type}, trial={self.trial}, threads={self.threads}, time={self.real_sec} secs, " + \
            f"min_delay={self.min_delay}, max_delay={self.max_delay}, inc_amnt={self.increment_amnt})"


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
    experiment = Experiment(experiment_trials, announce_findings = True)

    print("example experiment trial: ",experiment_trials[120])    

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

    # experiment.graph_speedup_trials_all(path_to_exp_data_dir)
    # experiment.graph_process_trials_all(path_to_exp_data_dir)

    experiment.graph_process_trials_all(path_to_exp_data_dir)

