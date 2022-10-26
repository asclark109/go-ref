# Module 1

---

### Architectures Overview

__Modern architectures resemble Von Neumann architecture__ (as opposed Harvard Architecture)

Two main architectures people talk about: __Von Nuemann__ and __Harvard Architecture__


| Characteristic   |     __Von Neumann__      |  __Harvard__ |
|----------|-------------|------|
| physical address |  one physical address for accessing, storing both data, instructions | two physical addresses for accessing and storing both data, instructions (physically separate addresses) |
| buses / signal paths |    uses a single (common) path (bus) for transfer of data, instructions   |   uses separate buses for transfer of data, instructions |
| cycles |    requires 2 clock cycles for executing instruction (fetch, execute)   |   only requires 1 clock cycle to execute instruction (fetch and execute) |
| cost |    cheap   |   more expensive |
| access to CPU |    CPU can't read/write data and access instructions at same time   |   CPU easily read/writes data and accesses instructions at same time |
| uses |    small computers, personal computers (general purepose PC)   |   signal processing, micro controllers (embedded systems) |
| Hardware requirements |    requires lesser architecture   |   requires more hardware |
| Space requirements |    requires less space (data + instructions stored in same place)   |   requires more space (data, instructions stored separately) |
| usage of space |    efficient   |   less efficient |
| execution speed |    slower   |   faster |
| controlling |    simpler   |   more complex |

![alt text](pics/architectures.JPG "Title")

---

### Basic Architecture (Von Neumann)

![alt text](pics/Computer_architecture_block_diagram.png "Title")

__Memory__: a collection of locations, capable of storing instructions or data
* Every location consists of an __address__, which is used to access the location, and the contents of the location.
* A computer can contain various types of memory components (registers, caches, RAM, secondary memory (i.e., hard drive) etc.)
  * registers --> secondary memory (faster access, less storage --> slower access, more storage)

__Central Processing Unit (CPU)__: controller of a computer and is composed of many different parts. Two main parts:
* __Control Unit (CU)__: responsible for deciding which instruction in a program should be executed. (the boss) 
* __Arithmetic and Logic Unit (ALU)__: responsible for executing the actual instructions. (the worker) 

__Input Devices__: devices that send information to a computer for processing (e.g., keyboard, mouse, etc.)

__Output Devices__: devices that display/use the results of processing of tasks by the computer. (e.g., monitor, printer, speakers) 

---

### Computer Program Basics

![alt text](pics/serialProblem.gif "Title")

__Sequential Program__: 
* series of textual statements specifying the order of execution
* each statement is converted to an instruction by CPU.
* programs may be deterministic (for given input, same ordered statements are executed) or non-deterministic
  * note: some people use the word deterministic to mean deterministic results (same results for a given input)

Speed of a single processor is driven by __transistor density__ on the CPU chip
* as the size of transistors goes down, their speed can be increased, leading to overal speed of the integrated chip increasing (i.e. increasing clock rate)

__Moore's Law__:
* computing power doubles every two years
* density of transistors approximately doubles every two years

~Three Problems are encroaching on progress!
1. __<span style="color:red">Heat</span>__. As transistor density goes up, power consumption per volume increases, heat generation per volume increases, air cooling no longer becomes sufficient to keep devices cool 
2. __<span style="color:red">Structural Integrity (fragile)</span>__. As transistor density goes up, transistor size is going down, transistor fragility goes up. 
3. __<span style="color:red">Production Issues</span>__. As transistor density goes up, quantum tunneling and other production issues occur, production speed goes down. 

Speed of integrated circuit is reaching a limit of ~3.0-3.7 GHz. With overclocking, 4.0-5.0 GHz.

---

### Multi-Core Chips

__Multi-Core Chip Design__: can continue to increase transistor density by having multiple and lower clocked processors on one CPU chip

__Core__: "processor". synonymous with CPU

__Mult-Core__: "multiprocessor". more than one core on integrated circuit

![alt text](pics/multicore.JPG "Title")
> ex: single-core processor. multi-core processor. multi-core processor.

Intel's newest processors (January 2022) used in i9-12900K core, i7-12700K core, ... have
* 8-6 performance CPU cores (higher clock rates); physically large
* 8-4 efficient CPU cores (lower clock rates); physically small

### Programming on Multi-cores

Programmer needs to break the problem into discrete _tasks_, where each task can run on separate core.

__Task__:
* a unit of computational work. simply put, work is a series of program statements.
* tasks can be run (potentially) in parallel, but this is not a requirement

__Parallel Program__:
* solves a problem that consists of multiple tasks running on multiple processors simultaneously

![alt text](pics/parallelProblem.gif "Title")
> e.g. we took 1 problem and broke it down into 4 tasks, where each task has a series of instructions

---

### Concurrency vs Parallelism

Both concepts refer to a notion of executing a task at the same time.

__Concurrency__: 
* dealing with tasks that are __logically__ happening simultaneously.
* Two tasks are concurrent if they can be logically active at the same point in time

__Parallelism__:
* dealing with tasks that are __physically__ happening simultaneously

_Concurrency_ provides a way to structure a solution to a problem that may (not necessarily) have tasks executed in parallel!

From: https://freecontent.manning.com/concurrency-vs-parallelism/

__Concurrency__:
* is about multiple tasks which start, run, and complete in overlapping time periods, in no specific order.

__Parallelism__:
* is about multiple tasks or subtasks of the same task that literally run at the same time on a hardware with multiple computing resources like multi-core processor. As you can see, concurrency and parallelism are similar but not identical.

__Concurrency__ is a semantic property of a program or system. Concurrency is when multiple tasks are in progress for overlapping periods of time. Note, here we are not talking about the actual execution of the tasks, but the design of the system – that the tasks are order-independent. So, concurrency is a conceptual property of a program or a system, it’s more about how the program or system has been designed.

Imagine that one cook is chopping salad while occasionally stirring the soup on the stove. He has to stop chopping, check the stove top, and then start chopping again, and repeat this process until everything is done.

![alt text](pics/chop_stir.JPG "Title")

As you can see, we only have one processing resource here, the chef, and his concurrency is mostly related to logistics; without concurrency, the chef has to wait until the soup on the stove is ready to chop the salad.

__Parallelism__ is an implementation property. Parallelism is literally the simultaneous physical execution of tasks at runtime, and it requires hardware with multiple computing resources. It resides on the hardware layer.

Back in the kitchen, now we have two chefs, one who can do stirring and one who can chop the salad. We’ve divided the work by having another processing resource, another chef.

![alt text](pics/chop_stir_par.JPG "Title")

Parallelism is a subclass of concurrency: before you can do several tasks at once, you have to manage several tasks first.

The essence of the relationship between concurrency and parallelism is that concurrent computations can be parallelized without changing the correctness of the result, but concurrency itself does not imply parallelism. Furthermore, parallelism does not imply concurrency; it is often possible for an optimizer to take programs with no semantic concurrency and break them down into parallel components via such techniques as pipeline processing, wide vector SIMD operations, or divide and conquer.

As Rob Pike pointed out “Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.” In a single-core CPU, you can have concurrency but not parallelism. But both go beyond the traditional sequential model in which things happen one at a time.

To get more idea about the distinction between concurrency and parallelism, consider the following points:

- _An application can be concurrent but not parallel_, which means that it processes more than one task at the same time, but no two tasks are executing at the same time instant.
- _An application can be parallel but not concurrent_, which means that it processes multiple sub-tasks of a single task at the same time.
- _An application can be neither parallel nor concurrent_, which means that it processes one task at a time, sequentially, and the task is never broken into subtasks.
- _An application can be both parallel and concurrent_, which means that it processes multiple tasks or subtasks of a single task concurrently at the same time (executing them in parallel)

Imagine you have a program that inserts values into a hash table. If you spread the insert operation between multiple cores, that’s parallelism. But coordinating access to the hash table is concurrency.

---

![alt text](pics/conc_v_par.JPG "Title")

__NOTE__: each parallel program should outperform the single processor concurrent and sequential program!

---

# Module 2

---

## General Parallel Computing Terminology

__CPU__: Contemporary CPUs consist of one or more cores - a distinct execution unit with its own instruction stream.
* Cores with a CPU may be organized into one or more sockets - each socket with its own distinct memory.
  * When a CPU consists of two or more sockets, usually hardware infrastructure supports memory sharing across sockets.

__Node__: A standalone "computer in a box." 
* Usually comprised of multiple CPUs/processors/cores, memory, network interfaces, etc.
* Nodes are networked together to comprise a supercomputer.

__Task__: A logically discrete section of computational work.
* A task is typically a program or program-like set of instructions that is executed by a processor.
* A parallel program consists of multiple tasks running on multiple processors.

__Pipelining__: Breaking a task into steps performed by different processor units, with inputs streaming through, much like an assembly line; a type of parallel computing.

__Shared Memory__: a computer architecture where all processors have direct access to common physical memory.
* In a programming sense, it describes a model where parallel tasks all have the same "picture" of memory and can directly address and access the same logical memory locations regardless of where the physical memory actually exists.

__Symmetric Multi-Processor (SMP)__: Shared memory hardware architecture where multiple processors share a single address space and have equal access to all resources - memory, disk, etc.

__Distributed Memory__: in hardware, refers to network based memory access for physical memory that is not common.
* As a programming model, tasks can only logically "see" local machine memory and must use communications to access memory on other machines where other tasks are executing.

__Communications__: Parallel tasks typically need to exchange data.
* There are several ways this can be accomplished, such as through a shared memory bus or over a network.

__Synchronization__: the coordination of parallel tasks in real time, very often associated with communications.
* Synchronization usually involves waiting by at least one task, and can therefore cause a parallel application's wall clock execution time to increase.

__Computational Granularity__: in parallel computing, granularity is a quantitative or qualitative measure of the ratio of computation to communication.
* __Coarse__: relatively large amounts of computational work are done between communication events
* __Fine__: relatively small amounts of computational work are done between communication events

__Observed Speedup__: Observed speedup of a code which has been parallelized, defined as:
```
        wall-clock time of serial execution
        -----------------------------------
        wall-clock time of parallel execution
```
One of the simplest and most widely used indicators for a parallel program's performance.

__Parallel Overhead__: required execution time that is unique to parallel tasks, as opposed to that for doing useful work.
Parallel overhead can include factors such as:
* Task start-up time
* Synchronizations
* Data communications
* Software overhead imposed by parallel languages, libraries, operating system, etc.
* Task termination time

__Massively Parallel__:
* Refers to the hardware that comprises a given parallel system - having many processing elements. The meaning of "many" keeps increasing, but currently, the largest parallel * computers are comprised of processing elements numbering in the hundreds of thousands to millions.

__Embarrassingly (IDEALY) Parallel__:
* Solving many similar, but independent tasks simultaneously; little to no need for coordination between the tasks.

__Scalability__: Refers to a parallel system's (hardware and/or software) ability to demonstrate a proportionate increase in parallel speedup with the addition of more resources. Factors that contribute to scalability include:
* Hardware - particularly memory-cpu bandwidths and network communication properties
* Application algorithm
* Parallel overhead related
* Characteristics of your specific application

## Potential Benefits, Limits and Costs of Parallel Programming

__Complexity__: 
* In general, parallel applications are more complex than corresponding serial applications. Not only do you have multiple instruction streams executing at the same time, but you also have data flowing between them.
* The costs of complexity are measured in programmer time in virtually every aspect of the software development cycle:
  * Design
  * Coding
  * Debugging
  * Tuning
  * Maintenance
* _Adhering to "good" software development practices is essential when developing  parallel applications._

__Portability__: 
* Thanks to standardization in several APIs, such as MPI, OpenMP and POSIX threads, portability issues with parallel programs are not as serious as in years  past. However...
* All of the usual portability issues associated with serial programs apply to parallel programs. For example, if you use vendor "enhancements" to Fortran, C or C++, portability will be a problem.
* Even though standards exist for several APIs, implementations will differ in a number of details, sometimes to the point of requiring code modifications in order to effect portability.
* Operating systems can play a key role in code portability issues.
* Hardware architectures are characteristically highly variable and can affect portability.

__Resource Requirements__:
* The primary intent of parallel programming is to decrease execution wall clock time, however in order to accomplish this, more CPU time is required. For example, a parallel code that runs in 1 hour on 8 processors actually uses 8 hours of CPU time.
* The amount of memory required can be greater for parallel codes than serial codes, due to the need to replicate data and for overheads associated with parallel support libraries and subsystems.
* For short running parallel programs, there can actually be a decrease in performance compared to a similar serial implementation. The overhead costs associated with setting up the parallel environment, task creation, communications and task termination can comprise a significant portion of the total execution time for short runs.

__Scalability__:
Two types of scaling based on time to solution: strong scaling and weak scaling.

* __Strong scaling (Amdahl)__: 
  * The total problem size stays fixed as more processors are added.
  * Goal is to run the same problem size faster
  * Perfect scaling means problem is solved in 1/P time (compared to serial)

* __Weak scaling (Gustafson):__
  * The problem size per processor stays fixed as more processors are added. The total problem size is proportional to the number of processors used.
  * Goal is to run larger problem in same amount of time
  * Perfect scaling means problem Px runs in same time as single processor run

* The ability of a parallel program's performance to scale is a result of a number of interrelated factors. Simply adding more processors is rarely the answer.

* The algorithm may have inherent limits to scalability. At some point, adding more resources causes performance to decrease. This is a common situation with many parallel applications.

* Hardware factors play a significant role in scalability. Examples:
  * Memory-cpu bus bandwidth on an SMP machine
  * Communications network bandwidth
  * Amount of memory available on any given machine or set of machines
  * Processor clock speed

* Parallel support libraries and subsystems software can limit scalability independent of your application.

---

## Parallel Computer Memory Architecture

---

## Shared Memory System

has at least one multi-core CPU that allows all processors to access memory as a single global address space ("global memory")
* Shared memory parallel computers vary widely, but generally have in common the ability for all processors to access all memory as global address space.
* __Multiple processors can operate independently but share the same memory resources.__
* __Changes in a memory location effected by one processor are visible to all other processors.__
* Historically, shared memory machines have been classified as __UMA__ and __NUMA__, based upon memory access times.

---

### Uniform Memory Access (UMA) System

each core has equal access and access times to global memory

--- 

### Symmetric Multiprocessing (SMP) System

* each SMP processor is identical
* all processors can access shared main memory at same speed
* controlled by single operating system instance treating all processors equally (no processor reserved for specific purpose)
* tightly coupled multiprocessor system
  * all processors can execute different programs (and with different data) in parallel
  * all processors shared common resources (e.g. memory, I/O decies, ...)
  * all processors processors on same system bus

![alt text](pics/SMP.JPG "Title")

* processors and memory linked by __System bus__: broadcast medium like ethernet
* processors and memory have __bus controller units__: in charge of sending and listening for messages on the bus (listening == __"snooping"__)

__<span style="color:blue">(+)</span> SMP easy to build; common__ \
__<span style="color:red">(+)</span> SMP not scalable to large number of processors because the bus becomes overloaded; SMP configurations do not scale past 64 processors__ 

__NUMA architecture__ addresses the scaling issues of SMP architecture. Consists of Nodes that have cores and memory close together. Nodes communicate with each other over a network.

--- 

### System Bus (for SMP architecture)

* also called __"interconnect"__
* is a finite resource shared among the processors
* performance limited by __memory bus bandwidth__
* processors can be delayed if others are consuming too much of the interconnect's bandwidth
* SMP configurations does not scale well past 64 processors

--- 

## Program Execution

The Operating System (OS) and hardware are responsible for managing execution of a program.

---

### Process

* has its own block of memory (registers, stack, ...)
* has program code translated into machine language
* has information about state of the process (e.g. program counter)
* has security information
* has descriptors of resources the OS has allocated to the process
* __Can't__ access the memory of another process

Each process when started will spawn on a single __thread__

---

### Threads

an entity within the process that can be scheduled for execution on a processor by the OS
* is responsible for executing the instructions of the task assigned to it
  * Task == a unit of computational work; i.e., a series of instructions from the program
* a sequential program has only one thread ("_main/primary/heavy thread_": executes all code in the sequential program)
* a concurrent program can __fork__ ("__spawn__") additional "light-weight" threads within a given process to execute tasks
  * all threads in process can be then executed in parallel on various processors
* hope: when one thread blocks because it is waiting on a resource, another will have work to do and can run
* __have their own stack where they can store local variables, function calls, etc, _but_ share the heap where dynamically allocated items are stored with other threads of the same process__
* some programming languages and OS provide a notion of __thread-local storage__: threads can store/retreive values independent of other threads


![alt text](pics/thread_processes.jpg "Title")

---

### Processes and Processors

OS gives the illusion that a single processor system is running multiple programs simultaneously

__Time slice__: alloted potion of time given to a process to run

__Context Switch__: when a processor runs a process for a while and then sets it aside to run another process
* after a process' time is up, it waits (blocks) until it has a turn again

A processor may set aside (__deschedule__) a process for a number of reasons:
* a memory request may take some time to satisfy
* a process has run long enough (i.e. reached an end to its time slice); thus, it's time for another process to begin its time slice
* when a process is descheduled it may resume execution on another processor

---

### Simultaneous Multi-threading (SMT) / Hyperthreading

use several threads to schedule instructions from different threads in the same cycle (if possible)
* helps increase usage of functional units of a processor more effectively
* hardware support for SMT is based on replication of the chip area used to store process state ("AS" == "architectural state")
  * e.g. integer ALU and floating-point ALU could perform an integer operation and floating-point operation in the same clock cycle

![alt text](pics/HyperThreading-2-1.png "Title")

---

### Logical / Hardware Processors

The processor appears to the OS and user programs as a set of __logical processors__ to which processes, threads, can be assigned for execution.
* processes, threads come from a single or several user programs
* the number of replications of the processor state (AS, see prior section) determines the number of logical processors
* logical processors are also called __"hardware threads"__

---

### Structuring Parallel Programs

Consider these factors when structuring / designing parallel programs:
* __Task decomposition__: how to divide up a problem into subproblems (i.e. tasks) to be executed concurrently
* __Distribution of tasks__: how to structure parallel programs to efficiently assign tasks to threads and how this affects the scheduling of threads by the OS
* __Synchronization__: how tasks/threads communicate/coordinate in order to obtain deterministic results
* __Data Dependencies__: where one task's output is required to be the input for another task
* __Scalability__: As program is given more compute resources (i.e. better parallel hardware) its performance also increases

---

### Task Decomposition and Granularity

__Granularity__: how much work (amount of work / amount of computation) performed by a task\
__Grain size__: quantitative measure of computation / communication 
  * `= "time to complete task" / "time needed to exchange data between processors for task"`
  * for now, best to have many small tasks evenly distributed across all processors

---

## Improving Data Access Performance

---

### Processors and Memory

processors and main memory are __far apart__.
* __Long time__ to read from memory, write to memory
* __Longer time__ for processor to verify that the value has been written to memory

takes many cycles to complete memory-based request.
> similar to writing / mailing a letter vs making a phone call.

__Memory access time has a large influence on program performance__. \
The objective of architecture trends over the years has been to __reduce memory access latency__.
* __Latency__: total time that elapses until a memory access operation has been completely terminated 

Memory latency issues are alleviated by implementing a hierarchy of components that store data. \
Ranges from very few registers to one or more levels of small, fast caches to relatively slow main memory

![alt text](pics/memory_hierarchy.png "Title")

---

### Cache

takes 100s of CPU cycle to access main memory (for the processor)

__Cache__: collection of memory locations that can be access in less time than some other memory locations
* typically located on the same chip as the processors

__Cache Line__: fixed-size block data that also contains metadata (e.g. tag, index)
* typically 64 or 128 bytes

__Cache gives the ability to take data in data from main memory and bring it closer to the CPU so that it does not always need to use the interconnect to go to main memory.__ 

__Processor can also update data in its cache (updating main memory later).__

![alt text](pics/caches.JPG "Title")

### Principle of Locality

Caches are effective because most programs display a high degree of locality
* __Locality__: accessing one location is followed by an access of a nearby location
  * __Spatial Locality__: accessing a nearby location
    * e.g. accessing `L[i]`, likely to access `L[i+1]`
  * __Temporal Locality__: accessing in the near future
    * e.g. if indexing an array, might cache whole segment of the array and access it later

### Cache Levels

__L1__: typically resides on chip; __1-2 cycles to access__; ~32 KB \
__L2__: may reside on or off chip; __10 cycles to access__; ~256 KB \
__L3__: normally off chip; __< 30 cycles to access__; ~8 MB

![alt text](pics/cache_levels.JPG "Title")
> e.g. Intel i7-9xx

### Cache Hit / Miss

When a processor attempts to read a value from a given memory address, it first checks the caches

__Cache Hit__: processor finds something it is looking for in one of its caches (L1-L3) \
__Cache Miss__: processor does not find something it is looking for in any of its caches (L1-L3)

Caches are expensive to build and therefore much smaller than main memory

When a chace is full, need to make room for new entry by evicting existing entry. Can:
1. discard an entry if it has not been modified
2. write it back to main memory if it has been modified

__Replacement policy__: determines which cache line to replace to make room for a new location
* usually a least-recently used heuristic

### Cache Coherence

When two processors have x cached, and one processor updates x, there needs to be some protocol to keep the other cache in sync (get the updated value).

_MESI protocol_ is the most commonly referenced protocol.

__MESI protocol__: provides 4 states that a cache line can be in:
* __Modified (M)__: modified, a processor has modified cached data, must write back to memory
* __Exclusive (E)__: not modified, only one processor has a copy of a main memory data in a cache line.
* __Shared (S)__: shared not modified; a piece of data from mainm memory may be in different states
* __Invalid (I)__: cache line contents not meaningful

__Write-Through Cache Scheme__: when cache updates data in its cache, it __immediately__ broadcasts changes __and__ writes to memory
* __<span style="color:blue">(+)</span>__ memory, caches always agree
* __<span style="color:blue">(+)</span>__ more read hits, maybe
* __<span style="color:red">(-)</span>__ bus traffic on all writes
* __<span style="color:red">(-)</span>__ most writes to unshared data (e.g. a procesor incrementing an index in a for loop)

__Write-Back Cache Scheme__: when cache updates data in its cache, it __immediately__ broadcasts changes. Caches mark the data as dirty. When the cache line is replaced by a new cache line from memory, the dirty line is written back to memory (__memory updated later__).
* dirty data is held in a __write-buffer__ that will eventually write its contents back to main memory periodically

---

### Race Conditions

Shared memory systems use __shared variables / shared resources__ (i.e. memory locations), which can be accessed by all processors. Communication and cooperation between the processors is organized by writing and reading shared variables that are stored in memory.

__Race Condition__: Non-deterministic behavior in a parallel program when the result of the operation depends on the interleaving of certain individual operations.

* code with a race condition can sometimes run deterministically (have deterministic results) and fail other times
* hard to reproduce and diagnose because they appear infrequently
* sometimes only appear under heavy load or when using certain compilers, platforms, architectures
* __Data Race Condition__: when at least two threads access a shared variable at the same time. One thread is trying to modify the variable.
  * _normally caused by not using proper synchronization when accessing the shared variable_
  * races can also happen with files and I/O (e.g. printing to screen)

![alt text](pics/data_race.JPG "Title")


To ensure determinism and to avoid race conditions, you must determine _critical sections_ in your code.

__Critical Section__: Block of code where potentially more than one thread can execute the code at the same time (and this is potentially where shared resources are accessed/modified).

- code in a critical section _should be executed serially (synchronized)_. Only one thread should be allowed to execute that code. I.e., one thread enters and finishes critical section, then another thread enters and leaves, ...
- eventually another thread should be able to access this section once one thread has completed the critical section.

---

### Synchronization using `atomics`

_Synchronization_ is needed when dependencies exist between parallel tasks and/or to handle race conditions. Many low-level synchronization primitives (e.g. locks, monitors, etc) are built off of specialized hardware primitives/instructions (__"Atomic" operations__)

__Atomic Operation__: operation that completes in a single step relative to other threads (in a shared memory system).

- No other thread can observe the modification to that shared variable halfway through its operation.

__Single-File Line Pattern__
```go
for !atomic.CompareAndSwapInt64(flag,0,1) {
}

// ...
// CRITICAL SECTION
// ...

atomic.StoreInt64(flag,0)
```

### Issues with Atomic operations

Best practice is to use atomic operations sparingly. They are not the best thing to be using.

* __<span style="color:red">(-)</span>__ Atomic operations are implemented using __CAS__ (compareAndSwap) or (__LL/SC__), which take significantly more clock cycle to complete than a simple load or store operation
* __<span style="color:red">(-)</span>__ atomic operations cause a __memory fence__, which forces the write-back buffer to be sent to main memory. This process can then stall other processes from reading/writing to main memory.
* __<span style="color:red">(-)</span>__ prevents out-of-order execution and various compiler optimizations.

cost-to-performance varies depending on architectures, program design, etc. \
These primitive atomic operations add more complexity to hardware.

---

### Instruction Level Parallelism (ILP)

For multi-core architectures, the basic unit of time is a cycle

__Cycle__: time it takes a processor to fetch and execute a single instruction (harvard architecture)
> 1980: 10 million cycles / second (0.01 Ghz)\
> 2005: 3  billlion cycles / second (3.0 Ghz)

Some primitive instructions take one cycle, and some take 100s of cycles

__Instruction-Level Parallelim (ILP)__: simultaneous execution of a sequence of instructions
* serial and parallel programs benefit from this in multicore modern hardware architectures

* __<span style="color:blue">(+)</span>__ __Instruction Pipeling__: uses functional units (e.g. ALUs, FPUs, load/store unbits, branch units) to execute independent instructions in parallel by different functional units
> e.g. the compiler can reorder instructions to keep the floating point unit (FPU) busy rather than be idle

![alt text](pics/pipelining.JPG "Title")

* __<span style="color:blue">(+)</span>__ __Out-of-order Execution__: instructions execute in any order that does not violate data dependencies

* __<span style="color:blue">(+)</span>__ __Speculative Execution and Branch Prediction__: processors can execute instructions speculatively before branches or data have been computed
> e.g. when an `if` statement starts, will begin to compute true and false branches before even evaluating the conditional

__Superscalar processor__: processor that allows for ILP

These techniques are handled by compilers and hardware to improve performance of sequential and parallel programs

---

# Module 3

---

## Admdahl's Law

![equation](http://www.sciweavers.org/tex2img.php?eq=%5Cbegin%7Bequation%7D%0A%5Cbegin%7Bsplit%7D%0At_%7Bsequential%7D%20%26%3A%3D%201%20%5C%5C%0Ap%20%26%3A%3D%20%5Ctext%7Bfraction%20of%20program%20that%20can%20be%20parallelized%7D%20%5C%5C%0An%20%26%3A%3D%20%5Ctext%7Bnumber%20of%20concurrent%20processors%7D%20%5C%5C%0At_%7Bparallel%7D%20%26%3D%201-p%20%2B%20%20%5Cfrac%7Bp%7D%7Bn%7D%20%5C%5C%5C%5C%5C%5C%0AS%20%26%3A%3D%20speed%20up%20%3D%20%5Cfrac%7Bt_%7Bsequential%7D%7D%7Bt_%7Bparallel%7D%7D%20%3D%20%5Cfrac%7B1%7D%7B%201-p%20%2B%20%20%5Cfrac%7Bp%7D%7Bn%7D%20%7D%0A%5Cend%7Bequation%7D%0A%5Cend%7Bsplit%7D&bc=White&fc=Black&im=jpg&fs=12&ff=arev&edit=0)

> example: can parallelize 40% of the code (work). What overall speed is achieved with 10 threads? p == 0.4. n == 10. 1 / (1-.4 + 0.4/10) = 1 / (0.6 + 0.04) = 1.56

---

## Locks

Use locks to create a critical section, a block of code that should be executed by only one thread at a time. Locks are built on top of atomic operations.

In Java,
```java
// Lock interface comes from java.util.concurrent.locks
Lock mutex = new LockImpl(...); // lock implementation
// ...
mutex.lock()
try {
  //...
} finally {
  //... restore invariant if needed
  mutex.unlock()
}
```

__Well-formed thread__
1. each critical section is associated with a `Lock` object
2. each thread calls that object's `lock()` method when it wants to enter the critical section
3. each thread calls `unlock()` when it leaves the critical section

- thread __acquires__ a lock when it returns from a `lock()` method call ("locks")
- thread __releases__ a lock when it invokes the `unlock()` method call ("unlocks")
- thread __holds__ a lock when it has acquired a lock but has not subsequently released it
- lock is __busy__ if a thread holds it
- lock is __free__ if no thread holds it

---

## Formal Properties in asynchronous computation

- __Safety Properties__: nothing bad ever happens
- __Liveness Properties__: states that a particular "good" thing will happen

## Formal Properties of a good Critical Section

Synchronization primitives need to adhere to the following properties and principles about critical sections in order to be correct:

* <span style="color:blue"> __Mutual Exclusion__</span> (_safety property_): only one thread is executing a critical section at a time. critical sections of different threads do not overlap. At most one thread holds the lock at any time.

<span style="color:green"> _Guarantees a computation's results are correct_.</span>

* <span style="color:blue"> __Deadlock-Freedom__</span> (_liveness property_):
  * If multiple threads simultaneously request to enter a critical section, then it must allow one to proceed.
  * Threads outside the critical section have no say in which thread can proceed into the critical section--only those currently waiting have influence.
  * If a thread is attempting to acquire/release a lock, then eventually some thread acquires/releases lock.
  * If a thread invokes `lock()` or `unlock()`, then it eventually returns from invoking `lock()` or `unlock()`.
  * <span style="color:green"> _If a thread is trying to enter critical section, then some thread, not necessarily the same one, eventually will enter critical section. OR At least one, always wins_.</span>

<span style="color:green"> _implies system never "freezes"_.</span>

* <span style="color:blue"> __Starvation-Freedom__</span> (_liveness property_)
  * Every thread that attempts to acquire (or to release) the lock eventually succeeds.
  * every call to `lock()`, `unlock()` eventually returns.

  * <span style="color:green"> _every thread trying to enter critical section, will eventually enter critical section._</span>

> __deadlock-freedom__ is saying that there are some threads will make progresses, but others might be stuck(starving), trying to get into critical section. Not all threads are stuck, so there is no deadlock, i.e. deadlock-freedom.

> On other hand, __starvation-freedom__ is saying that every process trying to get into critical section, will eventually do so. There will be no processes that will ever starve.

> This makes starvation-freedom a much stronger property than deadlock-freedom

> starvation-freedom implies deadlock-freedom

_many mutually exclusive algorithms in practice are not starvation free because starvation will be unlikely in those algorithms_

_there is no guarantee on how long a thread will wait to acquire the lock_

* <span style="color:blue"> __Fairness Principle__</span>: a thread that just left the critical section cannot immediately re-enter the critical section if other threads have already requested to enter the critical section.

_some algorithms bound how long a thread can wait_

---

### `LockOne` 2-thread lock algorithm (not good)

```java
class LockOne implements Lock {

  // volatile keyword here makes sure that
  // the changes made in one thread are 
  // immediately reflect in other thread
  private volatile boolean[] flag = new boolean[2];
  // thread-local index, 0 or 1

  public void lock() {
    int i = ThreadID.get();   // this thread == i
    int j = 1 - i;            // other thread == j
    flag[i] = true;           // this thread signals it wants the lock
    while (flag[j]) {}        // wait until flag[j] == false
    }

  public void unlock() {
    int i = ThreadID.get();
    flag[i] = false;          // this thread gives up lock
  }
}
```

> In practice, the Boolean flag variables, as well as the victim and label
variables in later algorithms, must all be declared volatile to work properly.

- __Sequential exectution__: okay!
- __Concurrent exectution__: can fail! Does not exhibit deadlock-freedom. Consider the case that each thread calls lock() and sets their flag to true. They will get stuck waiting for the other to give up their desire for the lock.
> A: flag[i] = true;     B: flag[i] = true; \
> A: while (flag[j]) {}  B: while (flag[j]) {} 


### `LockTwo` 2-thread lock algorithm (not good)

```java
class LockTwo implements Lock {

  private volatile int victim; // indicates who should yield
  
  public void lock() {
    int i = ThreadID.get();    // this thread == i
    victim = i;                // let the other go first; assign self to be the victim--to yield
    while (victim == i) {}     // wait
  }
  
  public void unlock() {}
}
```

- __Sequential exectution__: Fails! deadlocks. first thread will sit and wait forever until a new thread comes.
- __Concurrent exectution__: okay! if thread i in critical section, then victim == j, which implies mutual exclusion

### `PetersonLock` 2-thread lock algorithm (good)

```java
class Peterson implements Lock {

  // thread-local index, 0 or 1
  private volatile boolean[] flag = new boolean[2];
  private volatile int victim;

  public void lock() {
    int i = ThreadID.get();    // this thread == i
    int j = 1 - i;             // this thread == j
    flag[i] = true;            // I’m interested
    victim = i;                // you go first
    while (flag[j] && victim == i) {} // wait; while other thread interested, and you're going first, I'll wait
  }

  public void unlock() {
    int i = ThreadID.get();
    flag[i] = false;           // I’m not interested anymore
}
```

- __Sequential exectution__: good!
- __Concurrent exectution__: good!

Mutually exlusive, starvation-free, deadlock-free.

---

### LiveLock

See hand-written notes on live-locking.

- __Deadlock__: system enters a state from which there is no way for threads to make progress (e.g. `LockOne`, `LockTwo` can deadlock). _this is a narrow definition for deadlock that is sometimes used in literature or casually_.
- __Livelock__: two or more threads actively prevent each other from making progress by taking steps that subvert steps taken by other threads. _when a system is livelocked, there is a way to schedule threads so that the system can make progress (but also some way to schedule them so there is no progress_. _some people take a broader definition for deadlock to include this kind of phenomenon_).

### Fairness

Starvation-Freedom implies every thread calling `lock()` eventually enters critical section. It does not say anything about how long it will take or whether it is fair.

Can conceptually splits a `lock()` call into a _doorway_ section and a _waiting_ section
* __Doorway section__ always completes in a bounded number of steps
* __Waiting section__ may take an unbounded number of steps

__Bounded Wait-Free__: section of code guaranteed to complete in a bounded number of steps.

__First-Come-First-Served__: if the lock can be split into a doorway section and waiting section such that whenever thread A finishes its doorway before B starts its doorway, then A cannot be overtaken by B:

![equation2](http://www.sciweavers.org/tex2img.php?eq=%5Ctext%7Bif%20%7DD_A%5Ej%20%5Crightarrow%20D_B%5Ek%5Ctext%7B%2C%20then%20%20%7DCS_A%5Ej%20%5Crightarrow%20CS_B%5Ek&bc=White&fc=Black&im=jpg&fs=12&ff=arev&edit=0)

> where the right arrow denotes "precedes" and e.g. D_A^j represents the doorway interval (elapsed time) for thread A on the jth time of requesting the lock. CS_A^j represents the critical section interval (elapsed time) that thread A spends within the critical section on the jth time of entering the critical section.

__deadlock-free__ ^ __first-come-first-served__ implies __starvation-free__.

---

### Bakery Algorithm `BakeryLock` n-thread lock algorithm; `O(n)` space

```java
class Bakery implements Lock {

  volatile boolean[] flag; // may need to be volatile?
  volatile Label[] label;  // may need to be volatile?

  public Bakery (int n) {
    flag = new boolean[n]; // initially: [false, false, false, ...]
    label = new Label[n];  // initially: [0,     0,     0,   , ...]
    for (int i = 0; i < n; i++) {
      flag[i] = false; label[i] = 0;
    }
  }

  public void lock() {
    int i = ThreadID.get();     // this thread == i
    flag[i] = true;             // this thread signals it wants the lock
    label[i] = max(label[0], ...,label[n-1]) + 1;  // this thread grabs a "ticket number"
    while ((∃k != i)(flag[k] && (label[k],k) << (label[i],i))) {}; // while there is a thread k
                                                                   // that wants the lock and has
                                                                   // a lower number than me, I'll wait
  }

  public void unlock() {
    flag[ThreadID.get()] = false; // this thread signals it doesn't want lock anymore
  }
 }
```

- __mutually exclusive__: yes!
- __deadlock-free__: yes!
- __first-come-first-serve__: yes!
- __space__: `O(n)`; where `n` is the number of threads

This lock algorithm is great--except for that it uses lots of space (linear in the number of threads). Will also learn that the implementation above as written may not work.

__<span style="color:blue">(+)</span> simple, correct__ \
__<span style="color:blue">(+)</span> fair (first-come-first-serve)__ \
__<span style="color:red">(-)</span> space `O(N)`, where `N == threads`__

---

### Spin Locks, Contention

If there is a mutual exclusion protocol for locking, threadss will often not get a lock when trying to acquire it (they won't acquire it immediately). In this situation, the thread has two possibilities

* __Spin Lock__: a lock that when trying to acquire the lock will keep trying to acquire the lock.
  * __Spinning__: a thread is spinning if it is repeatedly testing the lock ("busy waiting"). e.g. Filter, Bakery algorithms have spin locks.
  * __<span style="color:blue">(+)</span> good when__ delays are expected to be short because switching to another thread is expensive.

* __Blocking Lock__: when a thread suspends itself and asks the operating system to schedule another thread on the processor.
  * __<span style="color:blue">(+)</span> good when__ delays are expected to be long because we can switch to another thread and get work done.

---

### Practical Lock Implementation

```java
class Peterson implements Lock {

  // thread-local index, 0 or 1
  private volatile boolean[] flag = new boolean[2];
  private volatile int victim;

  public void lock() {
    int i = ThreadID.get();    
    int j = 1 - i;             
    flag[i] = true;            // (1) compiler may swap execution of thisline with the line below it (line (2))
    victim = i;                // (2)
    while (flag[j] && victim == i) {} // compiler may replace this with a conditional statement that spins forever
                                      // if thread doesn't immediately enter critical section
  }

  public void unlock() {
    int i = ThreadID.get();
    flag[i] = false;
}
```
1. modern processors and programming languages do __not provide sequentially consistent memory__
2. modern processors and programming languages do __not guarantee program order amongs reads and writes by a given thread__
3. many multiprocessor architectures have writes to shared memory buffered into a special __write buffer__ ("store buffer"), where it is written to memory only when needed.

__To prevent reordering of operations resulting from write buffering__, modern architectures provide a special __memory barrier__ instruction (also called "__memory fence__") that forces outstanding operations to take effect
- synchronization methods include a memory barrier (e.g. `getAndSet()`, `compareAndSet()`, of `AtomicInteger` in Java, e.g. reads and writes to volatile fields in Java)
- in Java, declare object fields accessed by concurrent threas as __`volatile`__. This makes reads, writes atomic
- in Java, for compound operations that need to be atomic, use `java.util.concurrent.atomic` package (e.g. `AtomicReference<T>`, `AtomicInteger`)

---

### Test-And-Set Locks (`TASLock`) n-thread lock algorithm; O(1) space

`TASLock` uses `AtomicBoolean` (from `java.util.concurrent`), which implements atomic operations.

```java
public class AtomicBoolean {
  boolean value;

  public synchronized boolean getAndSet(boolean newValue) {
    boolean prior = value;
    value = newValue;
    return prior;
  }
}
```

```java
public class TASLock implements Lock {
  
  AtomicBoolean state = new AtomicBoolean(false);
  
  public void lock() {
    while (state.getAndSet(true)) {} // atomically replaces current value with true and returns prior value
                                     // if another thread holds the key, then state == true, and true replaces true, returning true,
                                     // so this thread keeps spinning
  }

  public void unlock() {
    state.set(false);
  }
}
```

__<span style="color:blue">(+)</span>__: `TASLock` is O(1) space\
__<span style="color:red">(-)</span>__: `TASLock` has poor scalability (bad performance with more threads); caused by bus traffic from getAndSet() calls that need exclusive copies of data

---

### Test-And-Test-And-Set Lock (`TTASLock`) n-thread lock algorithm; O(1) space

```java
public class TTASLock implements Lock {
  
  AtomicBoolean state = new AtomicBoolean(false);

  public void lock() {
    while (true) {
      while (state.get()) {};       // "LURKING": while some thread has the lock, wait (spin)
                                    // the thread with the key has unlocked. state == false now.
      if (!state.getAndSet(true))   // "POUNCING": spinning threads move to this line. roughly at the same time try to getAndSet()
                                    // the state one will set the state from false to true, returning the previous value false, 
                                    // which leads to a true if-statement. this thread enters this section and calls return,
                                    // entering the CS.
        return;
      }
    }

  public void unlock() {
    state.set(false);
  }
}
```

__<span style="color:blue">(+)</span>__: `TTASLock` is O(1) space\
__<span style="color:red">(-)</span>__: `TTASLock` has better scalability than `TASLock` (but still bad performance with more threads); bus traffic caused by thread unlock() call, which leads to _"invalidation storms"_

`TTASLock` performs MUCH BETTER than `TASLock`.

![alt text](pics/taslocks.JPG "Title")

* When a processor writes to memory, and has a copy of the data in cache, it must be an exclusive copy.
  * if it has a copy of the data in cache and has an exclusive copy, processor has a __cache hit__ and writes to its copy.
  * else, it does not have a copy in cache (it has a __cache miss__), and it requests a copy by broadcasting on the bus
    * a processor that has a copy will share it.
      * since the processor requesting the copy is going to write to it, all copies of that data on other processors becomes invalidated (so that the requestor has an exclusive copy when it writes)

__TASLock uses `getAndSet()`__
* a thread must request an exclusive copy of the `state` variable because it is intending to write (unless cache already has exclusive copy).
  * in such a case, all other threads have their copy of the data _invalidated_ because the thread may write to the variable
    * if multiple threads are spinning on the lock, almost every call to `getAndSet()` leads to a cache miss, and they must request on the bus to fetch the (unchanged) value.
    * Also, when the thread with the lock tries to release it, it may be delayed because the bus is monopolized by the spinners.

__TTASLock uses `get()` for threads spinning__
* intially lock is held by thread A
  * when thread B reads the lock the first time (`state`), it has a cache miss. It requests a shared copy since it is only reading it (does not need exclusive copy). Thread A and B now both have their copy of the data (`state`) in the shared state.
    * As thread A holds the lock, thread B spins, but it just reads the lock `state` already in its cache.

* However, when thread A lets go of the lock (sets lock state to false), it requests an exclusive copy because it is writing, which causes all spinning threads to invalidate their cached copies.
  * each cache has a cache miss, re-reads the new value, and they all (more or less) call `getAndSet()` to attempt to acquire the lock.
    * the first thread to succeed invalidates the others, which then re-read the value, causing a storm of bus traffic

In summary, `TTASLock` outperforms `TASLock` because `TTASLock` involves __local spinning__: where a thread repeatedly rereads cached values instead of repeatedly using the bus. With `TASLock`, spinning threads continuously request exclusive copies of the `state` variable, which cause invalidations in other caches, and a lot of bus traffic. The bus traffic in `TTASLock` occurs when the thread holding the lock unlocks, which invalidates the spinners' cached values of the `state` variable.

---

### Exponential Backoff Lock (`BackoffLock`) n-thread lock algorithm; O(1) space

Want to improve upon `TTASLock` by reducing the bus traffic

- want __reduced bus bandwidth used by spinning threads__
- want __reduced release/acquire latency__
- want __reduced acquire latency for an idle lock__

__Contention__: when multiple threads try to acquire the lock at the same time

Attempting to acquire a highly contended lock is bad: an attempt to acquire the lock contributes to bus traffic at a time when the thread's chances of acquire the lock are slim. Instead, it is more effective for threads to __back off__ for some duration.

IDEA: each time a thread fails to acquire lock, it doubles the expected back off time up to a fixed maximum.
- __Randomness ensures threads do not fall into a lockstep__
- __more fails imply higher-contention; hence, the doubling of the expected back off time__

```java
public class Backoff {

  final int minDelay, maxDelay;
  int limit;

  public Backoff(int min, int max) {
    minDelay = min;
    maxDelay = max;
    limit = minDelay;   // current delay limit
  }

  public void backoff() throws InterruptedException {
    int delay = ThreadLocalRandom.current().nextInt(limit);  // ThreadLocalRandom is random num generator for a thread
                                                             // delay is random int between [0,limit]
    limit = Math.min(maxDelay, 2 * limit);                   // double the limit for the next call (up to max value)
    Thread.sleep(delay);
  }

}
```

```java
public class BackoffLock implements Lock {

  private AtomicBoolean state = new AtomicBoolean(false);
  private static final int MIN_DELAY = 8;     // performant sensitive; sensitive to architecture?
  private static final int MAX_DELAY = 1024;  // performant sensitive; sensitive to architecture

  public void lock() {
    Backoff backoff = new Backoff(MIN_DELAY, MAX_DELAY); // create a new, internal Backoff obj for each lock() call
    while (true) {
      while (state.get()) {};
      if (!state.getAndSet(true)) {
        return;
      } else {
        backoff.backoff();
      }
    }
  }

  public void unlock() {
    state.set(false);
  }

  // ...

}
```

__<span style="color:blue">(+)</span>__: `BackoffLock` better than `TTASLock`\
__<span style="color:red">(-)</span>__: `BackoffLock` underutilizes critical section when lock is uncontrolled (i.e. spinning threads may be sleeping when lock is free)\
__<span style="color:red">(-)</span>__: `BackoffLock` may be unfair\
__<span style="color:blue">(+)</span>__: `BackoffLock`'s unfairness may actually help performance wise: granting repeated access to a thread allows it to cache data and have many cache hits.\
__<span style="color:red">(-)</span>__: `BackoffLock`' performance sensitive to parameters; not portable accross platforms\
__<span style="color:blue">(+)</span>__: `BackoffLock`' easy to implement


__GOOD__: easy to implement. better than `TTASLock`.\
__BAD__: sensitive to parameters; not portable accross platforms

---

### Queue Locks (`ALock`, `AndersonLock`) n-thread lock algorithm; O(N) space

IDEA: have threads wait in a _queue_:

- __<span style="color:blue">(+)</span>__: cache coherence traffic reduced because threads spin on a different location
- __<span style="color:blue">(+)</span>__:  better utilization of critical section: each thread notified by predecesor thread in queue
- __<span style="color:blue">(+)</span>__:  first-come-first-serve fairness

__Thread-Local Variable__: each thread has its own, independently initialized copy of the variable:
- it is not stored in shared memory
- it requires no synchronization
- it does not generate coherence traffic

```java
public class ALock implements Lock {
  ThreadLocal<Integer> mySlotIndex = new ThreadLocal<Integer> (){  // each thread has its own
    protected Integer initialValue() {
      return 0;
    }
  };

  AtomicInteger tail;       // the ALock has 1 tail
  volatile boolean[] flag;  // declared volatile so that no compiler optimatizations happen below: while(!flag[slot]) {};
  int size;

  public ALock(int capacity) {
    size = capacity;
    tail = new AtomicInteger(0);    // sharedMemory for threads: 0
    flag = new boolean[capacity];   // sharedMemory for threads: [false, false, ... , false]
    flag[0] = true;                 // denote 0th ticket number as first ticket number allowed to take lock
  }

  public void lock() {
    int slot = tail.getAndIncrement() % size;  // this thread grabs a ticket number (slot)...first thread gets 0...now, tail == 1
                                               // flag[slot] represents whether thread with slot number should go
    mySlotIndex.set(slot);   // this thread sets remembers its slot number
    while (!flag[slot]) {};  // while I am told that I am not to go, I'll wait.
  }

  public void unlock() {
    int slot = mySlotIndex.get();  // this thread gets its ticket number
    flag[slot] = false;            // this thread signals that it is no longer in CS
    flag[(slot + 1) % size] = true; // this thread informs next thread in line to go.
  }

}
```

__<span style="color:blue">(+)</span>__:  `ALock` better than `BackoffLock`; scalable\
__<span style="color:blue">(+)</span>__:  `ALock` easy to implement\
__<span style="color:blue">(+)</span>__:  `ALock` minimizes interval between when a lock is freed by one thread and acquired by another (reduced latency)\
__<span style="color:blue">(+)</span>__:  `ALock` guarantees no starvation\
__<span style="color:blue">(+)</span>__:  `ALock` is first-come-first-served\
__<span style="color:red">(-)</span>__:  `ALock` requires known bound `N` on the max number of concurrent threads\
__<span style="color:red">(-)</span>__: `ALock` has poor space efficiency. `O(N)` space per lock with `N` involved threads. Synchronizing L things requires `O(LN)` space\
__<span style="color:red">(-)</span>__: `ALock` may exhibit false-sharing (can be mitigated with padding)

Each thread spins on its locally cached copy of a single array location, greatly reducing traffic.\
_However_, contention may still occur because of _false sharing_.

__False Sharing__: occurs when adjacent data items (e.g. array elements) share a single cache line. A write to one item invalidates the whole cache line, which causes invalidation traffic to processors that are spinning on nearby unchanged items that happen to fall in the same cache line.

False Sharing can be _mitigated_ with __padding__: padding array elements so that distinct elements are mapped to distinct cache lines. Padding is easier in lower-level languages like C/C++, where programmer has control over the layout of objects in memory.

![alt text](pics/padding.JPG "Title")

---

### LL Locks (`CLHLock`) n-thread lock algorithm; O(L+N) space

`CLHLock` improves upon `ALock` by having less space overhead. This lock keeps a FIFO Linked-list queue of threads.

This kind of lock uses a `QNode` object, which denotes each thread's status. 
* `QNode.locked == true` indicates thread has either acquired the lock or is waiting for the lock
* `QNode.locked == false` indicates thread has released the lock

```java
class QNode {
  volatile boolean locked = false;
}
```

The Lock itself is represented as a virtual linked-list of `Qnode` objects (__linked-list FIFO queue is implicit__).
  * each thread refers to its predecessor through a thread-local `pred` variable
  * public `tail` field is an `AtomicReference<QNode>` to the node most recently added to the queue

```java
public class CLHLock implements Lock {
  AtomicReference<QNode> tail;    // each thread talks to the same tail (sharedMemory)
  ThreadLocal<QNode> myPred;      // each thread has its own myPred
  ThreadLocal<QNode> myNode;      // each thread has its own myNode

  public CLHLock() {
    tail = new AtomicReference<QNode>(new QNode()); // Lock initially sets tail to QNode w/ QNode.locked == false

    myNode = new ThreadLocal<QNode>() {  // each thread will create its own myNode
      protected QNode initialValue() {
        return new QNode();
      }
    };

    myPred = new ThreadLocal<QNode>() {  // each thread will create its own myPred
      protected QNode initialValue() {
        return null;
      }
    };
  }

  public void lock() {
    QNode qnode = myNode.get();           // this thread gains access to its own myNode
    qnode.locked = true;                  // this thread signals it wants the lock

                                          // currently, tail points to the predecessor of this QNode we've made
    QNode pred = tail.getAndSet(qnode);   // advance tail to point to this thread's QNode, and obtain a reference to 
                                          // what tail was pointing to before (pred)

    myPred.set(pred);                     // set this thread's "predecessor node" (myPred) to be pred

    while (pred.locked) {}                // spin on predecessor's QNode
  }

  public void unlock() {
    QNode qnode = myNode.get();           // this thread gains access to its own myNode
    qnode.locked = false;                 // sets its QNode locked state to false; this will cache invalidate the successor spinning on this QNode
    myNode.set(myPred.get());             // reuse predecessor node to be this thread's QNode.
  }

}
```

![alt text](pics/clh.JPG "Title")

__<span style="color:blue">(+)</span>__: `CLHLock` low space `O(L+N)` (`ALock` used `O(LN)` space); we can reuse already created QNodes\
__<span style="color:blue">(+)</span>__: `CLHLock` does not require upfront knowledge of the number of threads that might acquire lock (needed for `ALock`)\
__<span style="color:blue">(+)</span>__: `CLHLock` Lock release only affects successor [when a thread releases its lock, it invalidates only its successor's cache]\
__<span style="color:red">(-)</span>__: `CLHLock` does not work well on NUMA architecture (uncached); _memory is split across nodes in NUMA based architecture (communicate over network) thread will not be spinning on a localized copy of data_. _if there isn't caching in the NUMA architecture, there'll be lots of network traffic_. \
__<span style="color:red">(-)</span>__: `CLHLock` each thread spins on predecessor's memory; _making a thread point to another thread's node. so memory is farther away, so might cause some cache issues_

---

### LL Locks (`MCSLock`) n-thread lock algorithm; O(L+N) space

`MCSLock` improves upon `CLHLock` by having threads spin on their own QNodes. Better for NUMA architecture.

In `MCSLock`, the __linked-list FIFO queue is explicit__.

This kind of lock uses a `QNode` object __that also has a reference to its successor__
* `QNode.locked == true` indicates thread has either acquired the lock or is waiting for the lock
* `QNode.locked == false` indicates thread has released the lock
* `QNode.next` indicates the successor

```java
class QNode {
  volatile boolean locked = false;  // true indicates the thread is waiting for the lock
  volatile QNode next = null;
}
```

The Lock itself is represented as with an __explicit__ linked-list of `Qnode` objects.
  * each thread refers to its successor through `QNode.next`
  * public `tail` field is an `AtomicReference<QNode>` to the node most recently added to the queue

```java
public class MCSLock implements Lock {

  AtomicReference<QNode> tail;    // each thread talks to the same tail (sharedMemory)
  ThreadLocal<QNode> myNode;      // each thread has its own myNode

  public MCSLock() {
    
    tail = new AtomicReference<QNode>(null); // Lock initially sets tail to null
    
    myNode = new ThreadLocal<QNode>() {      // each thread will create its own myNode
      protected QNode initialValue() {
        return new QNode();
      }
    };

  }

  public void lock() {

    QNode qnode = myNode.get();           // this thread gains access to its own myNode

                                          // currently, tail points to the predecessor of this QNode we've made
    QNode pred = tail.getAndSet(qnode);   // advance tail to point to this thread's QNode, and obtain a reference to 
                                          // what tail was pointing to before (pred)

    if (pred != null) {                   // if there is a thread waiting to acquire lock or has the lock
      qnode.locked = true;                // this thread signals it wants (is waiting for) the lock
      pred.next = qnode;                  // updates the predecessor node's next field to point to this thread's QNode
      while (qnode.locked) {}             // this thread spins on its own QNode (waits for predecessor to update its successor)
    }

  }

  public void unlock() {

    QNode qnode = myNode.get();               // this thread gains access to its own myNode

    if (qnode.next == null) {                 // if this QNode's sucessor not defined: either (1) no successor or (2) slow updating successor
      if (tail.compareAndSet(qnode, null))    // returns true if tail is same as this thread's QNode (i.e. no successor)
                                              // if they are same, will set tail to be null.
        return;
      while (qnode.next == null) {}           // else there is a successor: wait until successor fills in its next field
    }

    qnode.next.locked = false;                // notify successor: set next.locked to false to indicate it is no longer waiting (lock is now free)
    qnode.next = null;
  }

}
```

__<span style="color:blue">(+)</span>__: `MCSLock` same advantages as `CLHLock`\
__<span style="color:blue">(+)</span>__: `MCSLock` does work well on cachless NUMA architectures because each thread controls the location on which it spins \
__<span style="color:red">(-)</span>__: `MCSLock` releasing a lock requires spinning (drawback) \
__<span style="color:red">(-)</span>__: `MCSLock` requires more reads, writes, and `compareAndSet()` calls than `CLHLock`

---

# Module 4

---

## Concurrent Data Structures

Harder to design concurrent data structures than sequential data structures.

Two issues arise in particular: __correctness__, __scalability__

Performance issues in lock-based systems:

- __Sequential Bottleneck__: at any time, at most one lock-protected operation doing any useful work
- __Memory Contention__: overhead traffic as a result of multiple threadas concurrently attempting to access the same memory location
- __Blocking__: if thread that currently holds the lock is delayed, then all other threads attempting to access are also delayed

An algorithm becomes called a __blocking algorithm__ when its implementation uses locks. Consider trying to make non-blocking (lock-free) algorithms

### Non-blocking Algorithms

ultimately implemented as a hardware operation (atomically combines a load and store)

> e.g. compare and swap (CAS)

__Lock-free Algorithm__:
* if there is guaranteed system-wide progress.
* while a given thread might be blocked by other threads, all CPUs can continue doing work without stalls

__Wait-free Algorithm__:
* if there is also guaranteed per-thread progress.
* all CPUs can continue doing work without stalls; but also, __no computation can ever be blocked by another computation__

A method is __lock-free__ if _some_ call always finishes in a finite number of steps\
A method is __wait-free__ if _every_ call finishes in a finite number of steps

## Coarse-grained synchronization

__Coarse-grained synchronization__:
* take a sequential implementation of a class
* add a scalable `lock` field
* ensure each method call acquires and releases the lock

---

__<span style="color:blue">(+)</span> works when levels of concurrency are low__ \
__<span style="color:red">(-)</span> when many threads try to access the object at the same time then the object becomes a sequential bottleneck__ \
__<span style="color:blue">(+)</span> simple, correct, understandable__ \
__<span style="color:red">(-)</span> works poorly with contention__

---


## Linked-List (concurrent)

```java
public interface Set<T> {
  boolean add(T x);
  boolean remove(T x);
  boolean contains(T x);
}
```

```java
private class Node {
  T item;
  int key;
  Node next;
}
```

### Linked-List (coarse-grained)

```java
public class CoarseList<T> {

  private Node head;
  private Lock lock = new ReentrantLock(); // This lock is reentrant: A thread that is holding the lock can acquire it again without blocking

  public CoarseList() {
    head = new Node(Integer.MIN_VALUE);
    head.next = new Node(Integer.MAX_VALUE);
  }

  public boolean add(T item) {

    Node pred, curr;
    int key = item.hashCode();

    lock.lock(); // LOCK DOWN LINKED-LIST

    try {

      pred = head;
      curr = pred.next;
      while (curr.key < key) {  // keep advancing until we find place to insert item (curr.key >= key)
        pred = curr;
        curr = curr.next;
      }

      if (key == curr.key) {    // item already in LL, return false
        return false;
      } else {                  // item not in LL, add it and return true
        Node node = new Node(item);
        node.next = curr;
        pred.next = node;
        return true;
      }
    } finally {
      lock.unlock(); // UNLOCK LINKED-LIST
    }
  }

  public boolean remove(T item) {

    Node pred, curr;
    int key = item.hashCode();

    lock.lock(); // LOCK DOWN LINKED-LIST

    try {

      pred = head;
      curr = pred.next;
      while (curr.key < key) { // keep advancing until we find place item would be to remove (curr.key >= key)
        pred = curr;
        curr = curr.next;
      }

      if (key == curr.key) {   // item in LL, remove it and return true
        pred.next = curr.next;
        return true;
      } else {                 // item not in LL, return false
        return false;
      }
    } finally {
      lock.unlock(); // UNLOCK LINKED-LIST
    }
  }
}
```

## Fine-grained synchronization

__Fine-grained synchronization__:
* partition the object into independently synchronized components
* allows method calls that access disjoint components to execute concurrently
  * methods will conflict when they access the same component at the same time

IDEA: two threads in separate sections of the data structure should be allowed to work concurrently

__Hand-over-hand locking__
`(-inf) -> (a) -> (b) -> (c) -> (d) -> (inf)`
1. lock `(-inf)`; lock`(a)`
2. unlock `(-inf)`; lock `(b)`
3. unlock `(a)`; lock `(c)`

---

__<span style="color:blue">(+)</span> threads can traverse linked-list in parallel__ \
__<span style="color:red">(-)</span> long chain of acquire/release of locks__ \
__<span style="color:red">(-)</span> inefficient__

---



### Linked-List (fine-grained)

```java
public class FineList<T> {

  private Node head;
  private Lock lock = new ReentrantLock(); // This lock is reentrant: A thread that is holding the lock can acquire it again without blocking

  public FineList() {
    head = new Node(Integer.MIN_VALUE);
    head.next = new Node(Integer.MAX_VALUE);
  }


  public boolean add(T item) {

    int key = item.hashCode();

    head.lock();        // lock first node
    Node pred = head;
    try {
      Node curr = pred.next;
      curr.lock();      // lock next node
      try {

        while (curr.key < key) {    // PERFORM hand-over-hand locking until curr.key (2nd node) >= key 
          pred.unlock();            // unlock first node
          pred = curr;
          curr = curr.next;
          curr.lock();              // lock third node
        }

        if (curr.key == key) {      // CASE: item already in linked-list; return false
          return false;
        }

        Node node = new Node(item); // CASE: item not in linked-list; add it and return true
        node.next = curr;
        pred.next = node;
        return true;

      } finally {
        curr.unlock();
      }

    } finally {
      pred.unlock();
    }
  }

  public boolean remove(T item) {

    int key = item.hashCode();

    head.lock();        // lock first node
    Node pred = head;

    try {
      Node curr = pred.next;
      curr.lock();      // lock next node
      try {

        while (curr.key < key) {    // PERFORM hand-over-hand locking until curr.key (2nd node) >= key 
          pred.unlock();            // unlock first node
          pred = curr;
          curr = curr.next;
          curr.lock();              // lock third node
        }

        if (curr.key == key) {      // CASE: item in linked-list; remove it and return true
          pred.next = curr.next;
          return true;
        }

        return false;               // CASE: item not in linked-list; return false

      } finally {
        curr.unlock();
      }

    } finally {
      pred.unlock();
    }
  }

}
```

![alt text](pics/handoverhand.JPG "Title")

- `Fig 9.8` shows that concurrent remove calls (`remove(a)`, `remove(b)`) can lead to incorrectness. `a` gets removed, but `b` does not.
- `Fig 9.9` shows how hand-over-hand solves the issue (locking 2 items)
- `Fig 9.10` shows why it is important that method calls acquire the locks in the same order (left to right); deadlock can occur otherwise (`add(a)`, `remove(b)`)

## Optimistic synchronization

__Optimistic synchronization__:
* e.g. in trees or lists, methods involve finding a component
  * if method finds sought-after component, it locks the component, and then checks the component has not changed in the interval between when it was inspected and when it was locked. If it has changed, it makes another attempt; otherwise, it proceeds

1. find nodes without locking
2. lock nodes
3. check everything is okay
---

__<span style="color:blue">(+)</span> limited-hotspots: `add()`, `remove()`, `contains()`. no contention on traversal__ \
__<span style="color:blue">(+)</span> wait-free traversal__ \
__<span style="color:blue">(+)</span> not a lot of lock acquire/release (only at destination)__ \
__<span style="color:red">(-)</span> need to traverse list twice (validation)__ \
__<span style="color:red">(-)</span> `contains()` method still acquires lock (~90% of calls in applications are `contains()`)__ \

---

Optimistic Linked-list implementation effective when: \
`(cost of scanning 2x w/o locks) << (cost of scanning 1x w/ locks)`


```java

public class OptimisticList<T> {

  private Node head;
  private Lock lock = new ReentrantLock(); // This lock is reentrant: A thread that is holding the lock can acquire it again without blocking

  public OptimisticList() {
    head = new Node(Integer.MIN_VALUE);
    head.next = new Node(Integer.MAX_VALUE);
  }


  public boolean add(T item) {
    int key = item.hashCode();

    while (true) {      // continuously attempt add() operation

      Node pred = head;
      Node curr = pred.next;
      while (curr.key < key) {          // scan without locking to find place to add item (curr.key >= key)
        pred = curr; curr = curr.next;
      }
                                        // lock down region in linked-list where we want to add item
      pred.lock();                      // lock pred
      try {
        curr.lock();                    // lock curr
        try {

          if (validate(pred, curr)) {   // CONFIRM pred still in LL and that pred points to curr

            if (curr.key == key) {      // passed validation: case: item already in LL; return false
              return false;
            } else {                    // passed validation: case: item not in LL; insert Node and return true
              Node node = new Node(item);
              node.next = curr;
              pred.next = node;
              return true;
            }
          }                             // validation failed! restart (back to while true).
        } finally {
          curr.unlock();
        }
      } finally {
        pred.unlock();
      }
    }
  }

  public boolean remove(T item) {
    int key = item.hashCode();

    while (true) {      // continuously attempt remove() operation

      Node pred = head;
      Node curr = pred.next;
      while (curr.key < key) {          // scan without locking to find place to remove item (curr.key >= key)
        pred = curr; curr = curr.next;
      }
                                        // lock down region in linked-list where we want to add item
      pred.lock();                      // lock pred
      try {
        curr.lock();                    // lock curr
        try {

          if (validate(pred, curr)) {   // CONFIRM pred still in LL and that pred points to curr

            if (curr.key == key) {      // passed validation: case: item in LL; remove and return true
              pred.next = curr.next;
              return true;
            } else {                    // passed validation: case: item not in LL; return false
              return false;
            }
          }                             // validation failed! restart (back to while true).
        } finally {
          curr.unlock();
        }
      } finally {
        pred.unlock();
      }
    }
  }

  public boolean contains(T item) {
    int key = item.hashCode();

    while (true) {      // continuously attempt contains() operation

      Node pred = head;
      Node curr = pred.next;
      while (curr.key < key) {          // scan without locking to find place to remove item (curr.key >= key)
        pred = curr; curr = curr.next;
      }
                                        // lock down region in linked-list where we are looking for item
      pred.lock();                      // lock pred
      try {
        curr.lock();                    // lock curr
        try {
          if (validate(pred, curr)) {   // CONFIRM pred still in LL and that pred points to curr
            return (curr.key == key);   // return whether item in LL
          }                             // validation failed! restart (back to while true).
        } finally {
          curr.unlock();
        }
      } finally {
        pred.unlock();
      }
    }
  }

  private boolean validate(Node pred, Node curr) {
    Node node = head;
    while (node.key <= pred.key) {
      if (node == pred)                // walk from head of LL until we find pred. (1) confirm pred accessible from head
        return pred.next == curr;      // if we found pred confirm also pred points to curr. (2) confirm pred points to curr as successor
      node = node.next;
    }
    return false;                      // otherwise either pred not in LL or pred doesn't point to curr as successor
  }
  
}
```

![alt text](pics/whyvalidation.JPG "Title")

## Condition Variables

__Condition Variable__: data object that allows a thread to suspend execution until a certain event or condition occurs
* when the event or condition occurs another thread can signal to the sleeping thread to "wake up"
* condition variable is _always_ associated with a mutex
  * __Monitor__: combined use of methods, mutual exclusion locks, and condition objects

### in go

`sync.Cond` package represents conditional variables

create a condition with `cond := NewCond(I Locker) *Cond`

```go
// Suspends calling thread and release the monitor lock
// when it resumes, it re-acquires the lock
// typically called when condition not true
func (c *Cond) Wait() {}

// resumes one thread waiting in Wait() if any
// called when condition becomes true and want to wake up one waiting thread
func (c *Cond) Signal() {}

// resumes all threads waiting in Wait()
// called when condition becomes true and wants to wake up all waiting threads
func (c *Cond) Broadcast() {}
```

example usage
```go
// Simple Example of Using a
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type SharedContext struct {
	mutex       *sync.Mutex
	cond        *sync.Cond
	wgContext   *sync.WaitGroup
	counter     int
	threadCount int
}

func worker(goID int, ctx *SharedContext) {

	//All threads compute the fib of 20
	fib(rand.Int() % 20)

	/******* barrier *******/
	ctx.mutex.Lock()
	ctx.counter++
  // some code here blah blah blah
	fmt.Printf("Goroutine #%v finished calculating fib(randNum).\n", goID)
	if ctx.counter == ctx.threadCount {
		ctx.cond.Broadcast()
	} else {
		for ctx.counter != ctx.threadCount {
			ctx.cond.Wait()
		}
	}
	ctx.mutex.Unlock()
	/*********************/

	if goID == 0 {
		fmt.Printf("Every finished. Yes!\n")
	}

	ctx.wgContext.Done()
}
func fib(x int) int {

	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)

}
func main() {

	var wg sync.WaitGroup
	threadCount := 10

	//Setup my go routine context with barrier for mutex and condition variable
	var mutex sync.Mutex
	condVar := sync.NewCond(&mutex)
	context := SharedContext{wgContext: &wg, threadCount: threadCount, cond: condVar, mutex: &mutex}

	//Spawn my threads to begin calculating fib number
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go worker(i, &context)
	}
	wg.Wait()
	fmt.Printf("Done.\n")
}
```

### in java

```java
// java lock interface
public interface Lock {
  void lock(); // block the caller until it acquires the lock
  void lockInterruptibly() throws InterruptedException; // like lock(), but throws an exception if the thread is interrupted while it is waiting**
  boolean tryLock();
  boolean tryLock(long time, TimeUnit unit);
  Condition newCondition(); // factory that creates and returns a Condition object associated with the lock
  // A condition is associated with a lock, and is created by calling that lock’s newCondition() method
  void unlock(); // releases the lock.
}

//// example usage of lock
mutex.lock();
try {
  queue.enq(x)
} finally {
  mutex.unlock();
}
////

// **Threads in Java can be interrupted by other threads.
// If a thread is interrupted during a call to a Condition’s await() method, the call throws InterruptedException.
// The proper response to an interrupt is application-dependent. see ex below for schematic example.
// To avoid clutter, we usually omit InterruptedException handlers from example code, even though they would be 
// required in actual code. (It is bad programming practice to ignore interrupts.)

//// example usage of condition
Condition condition = mutex.newCondition();

mutex.lock()
try {
  while (!property) { // not happy
    condition.await(); // wait for property
  } catch (InterruptedException e) {
    // ... application-dependent response
  }
  // happy: property must hold
}
////

public interface Condition {
  void await() throws InterruptedException; // await() and its variants release the lock, and give up the processor and then later awaken and reacquire the lock
  boolean await(long time, TimeUnit unit) throws InterruptedException;
  boolean awaitUntil(Date deadline) throws InterruptedException;
  long awaitNanos(long nanosTimeout) throws InterruptedException;
  void awaitUninterruptibly();
  void signal(); // wake up one waiting thread
  void signalAll(); // wake up all waiting threads
}
```

implementing a FIFO queue:
```java

// The LockedQueue class: a FIFO queue using locks and conditions. 
// There are two condition fields:
//   one to detect when the queue becomes nonempty,
//   and one to detect when it becomes nonfull.

class LockedQueue<T> {
  
  final Lock lock = new ReentrantLock();
  final Condition notFull = lock.newCondition();
  final Condition notEmpty = lock.newCondition();
  final T[] items;
  int tail, head, count;

  public LockedQueue(int capacity) {
    items = (T[])new Object[capacity];
  }

  public void enq(T x) {

    lock.lock();    // lock down queue
    try {
      while (count == items.length)  // go to sleep if queue full (give up lock)
        notFull.await();
      items[tail] = x;               // thread awoke (thread has lock): now add item

      if (++tail == items.length)    // wrap tail back to 0 idx
        tail = 0;

      ++count;
      notEmpty.signal();             // signal to wake up one thread to handle new task

    } finally {
      lock.unlock();                 // release lock
    }
  }

  public T deq() {

    lock.lock();    // lock down queue
    try {
      while (count == 0)
        notEmpty.await();            // go to sleep if queue empty (give up lock)

      T x = items[head];             // get item off queue
      if (++head == items.length)
      head = 0;
      --count;

      notFull.signal();              // signal to wake up one thread to handle vacant space in queue

      return x;

    } finally {
      lock.unlock();
    }
  }

}
```

## Semaphore

The semaphore is a generalization of the mutual exclusion lock.\
_motivation_: want to control access to a shared resource
> ex: system can only handle ~10 users at once; so you want to throttle logins to 10 max

__Semaphore__: 
* synchronization primitive used to control access to a shared resource by multiple threads
* has capacity `c`, which indicates how many threads can be in critical section at any time
  * semaphore is really a counter; it keeps track of the number of threads who've been granted permission to enter

__Binary Semaphore__: acts like a mutex; has capacity of 1

__Counting Semaphore__: capacity == number of available resources


### in go

For an implementation in go, we have capacity represent transient capacity; that is if semaphore capacity starts at 3, then when three want to enter the CS, each will decrement the semaphore by 1, bringing the "_capacity_" to 0. When a fourth thread wants to enter the CS, and tries to decrement the semaphore, it will go to sleep. When a thread leaves the CS it will increment the semaphore, waking up another thread.


```go
// example binary semaphore
var mutex_sema *Semaphore
mutex_sema = NewSemaphore(1)  // capacity of 1
mutex_sema.Down() // decrement semaphore (thread about to enter CS)
// critical section code
mutex_sema.Up()   // increment semaphore (thread left CS)

// example counting semaphore
var sema *Semaphore = NewSemaphore(10) // will restrict 10 threads to entering CS
```

implementation
```go
func (s *Semaphore) Down() {
  // wait untill semaphore is greater than 0
  // decrement value of semaphore by 1
}

func (s *Semaphore) Up() {
  // increment value of semaphore by 1
  // if there are 1 or more threads waiting, wake one up
}
```

example full implementation
```go
package semaphore

import "sync"

type Semaphore struct {
	capacity  int
	mutex     *sync.Mutex
	condition *sync.Cond
}

func NewSemaphore(capacity int) *Semaphore {
	var mutex sync.Mutex
	condition := sync.NewCond(&mutex)
	return &Semaphore{capacity, &mutex, condition}
}

func (s *Semaphore) Up() {
	s.mutex.Lock()
	s.capacity++
	s.condition.Broadcast()  // perhaps should be s.condition.Signal() call?
	s.mutex.Unlock()
}

func (s *Semaphore) Down() {
	s.mutex.Lock()
	for s.capacity == 0 {
		s.condition.Wait()
	}
	s.capacity--
	s.mutex.Unlock()
}

// "bounded buffer" Work Queue for Producer, Consumer Problem

```

### in java

This implementation keeps a `state int` and `capacity int` and manipulates the state variable. When thread wants to enter the CS, it calls `sema.acquire()`; the thread will go to sleep if `state` is at `capacity`.

```java
public class Semaphore {
  final int capacity;
  int state;
  Lock lock;
  Condition condition;

  public Semaphore(int c) {
    capacity = c;
    state = 0;
    lock = new ReentrantLock();
    condition = lock.newCondition();
  }
  
  public void acquire() {
    lock.lock();
    try {
      while (state == capacity) {
        condition.await();
      }
      state++;
    } finally {
      lock.unlock();
    }
  }

  public void release() {
    lock.lock();
    try {
      state--;
      condition.signalAll();
    } finally {
      lock.unlock();
    }
  }

}
```