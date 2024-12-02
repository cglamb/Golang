# Random Forest Analysis in Golang

Note:  Data is zipped and must be unzipped manually.  Data should be saved to a folder called Data 

**Author:** Charles Lamb  
**Contact Info:** charlamb@gmail.com  
**GitHub Project URL:** [Go Random Forest](https://github.com/cglamb/Go_Random_Forest)  
**Git Clone Command:** `git clone https://github.com/cglamb/Go_Random_Forest.git`

## Introduction
This project evaluates the performance of a random forest model across three programming languages: Golang, R, and Python. It aims to compare runtime and resource utilization using identical datasets, tree counts, and features considered at each split. This work extends my previous research on voting ensembles found [here](https://github.com/cglamb/Voting_Ensemble_Spam). EDA, data scrubbing, text preprocessing, and tokenization work is borrowed from that earlier work. The original data, a labeled SMS spam dataset from UCI Machine Learning and Esther Kim, is available on [Kaggle](https://www.kaggle.com/datasets/uciml/sms-spam-collection-dataset?resource=download).

## Experiment Design
All data processing is performed in Data_Prep.ipynb.  Test and training datasets are generated via this same Juypter notebook.  50% of original dataset was assigned to the training dataset and 50% to the testing dataset.  Both sets are exported as CSVs (and saved in this repository within a zip folder).  These CSV files were read into the Go, Python, and R codes.  By performing a single test/train split in Juypter notebook, we ensure all three software packages get the exact same test and training data.

## Overall Findings
Five separate models were fit in each language using a different combination of number of trees and max features at branches.  The exhibit below shows runtime and F-score for each model across each of the three languages.  Additional profiling reports are saved for both the Go and R scrips in the log folder.  These include metrics associated with memory usage.  

| # Trees | Features at Split | Go RunTime (secs) | Go Test F-Score | R RunTime (secs) | R Test F-Score | Python RunTime (secs) | Python Test F-Score |
|---------|-------------------|-------------------|-----------------|------------------|----------------|-----------------------|---------------------|
| 10      | 10                | 4.84              | 0.70            | 57.19            | 0.96           | 2.16                  | 0.73                |
| 100     | 10                | 23.82             | 0.74            | 377.31           | 0.96           | 2.46                  | 0.77                |
| 100     | 50                | 79.75             | 0.85            | 421.15           | 0.98           | 2.57                  | 0.83                |
| 100     | 100               | 128.96            | 0.86            | 374.77           | 0.98           | 2.71                  | 0.84                |
| 1000    | 10                | 299.60            | 0.73            | 3608.73          | 0.96           | 6.89                  | 0.84                |

*Note: Some metrics were not measured.

It should be noted that as default parametrization is different across the different languages and because random seeding was not controlled across the three different languages, the Go, R, and Python fitted models vary.  Although with similar F-scores on identical test data (at least between Go and Python), I am comfortable that the models are materially the same (and therefore the above runtime comparisons are more or less apples-to-apples).  It is interesting to note that the R models had a materially better F-score versus the other two models.  This bears further review but had not been researched further as of this writting.

## Recommendation
Go runtimes were extremely better than R.  On the other hand, Python experienced better runtimes than Go across every parameterization.  Additionally, the Python model appeared to scale better with runtime increasing by 2.8 times between 100 trees and 1000 trees (holding constant features at split), while Go had a 12.6 fold increase across the same dimension.

Go’s advantage versus R is not surprising as R’s random forest library computed serially, which meant Go’s ability to take advantage of concurrency resulted in substantially improved runtime.  On the other hand, Python’s substantial outperformance of Go is perhaps somewhat surprising.  Although the Python library used to perform the computation allows for parallel computation and thus eliminates most of Go’s advantage.

I suspect the underperformance of Go versus Python has to do with inefficiencies in the underlying random forest library used in Go.  The Go library used is a repository with 3 contributors and 114 stars, while the sklearn library used in Python is supported by a massive community with support from institutional and private grants.

My final recommendation is that Go is likely computational more efficient than R, and users in R should consider Go as an alternative…particularly if users are paying for compute (as in a cloud environment).  On the other hand, I do not see the same benefit versus Python…and my research suggests that users worried about computational efficiency should perhaps consider Python over Go.  


## Golang Implementation
The random forest was implemented in Golang using Andy Sun’s random forest library available at: https://github.com/fxsjy/RF.go/tree/master.  The library by default includes support for goroutines allowing for trees to be calculated concurrently.  Sun’s library also includes a prediction function which was used to make predictions against the test data.  Custom functions were developed to calculate accuracy, precision, recall, and f1 score.  Profiling was completed using the pprof library.  The resource utilization diagram produced by pprof is available at Log/Go/Go Profile (100 trees and 100 features at split)

## R Implementation
The random forest was implemented in R using Breiman and Cutler’s random forest library: https://www.stat.berkeley.edu/~breiman/RandomForests/.  Runtime and resource utilization was tracked using R Studio's (2023.12.1) built in profiling tool.  Profile reports and console logs are saved in /logs/R.

## Python Implementation
The random forest was implemented in Python using sklearn (https://scikit-learn.org/stable/modules/generated/sklearn.ensemble.RandomForestClassifier.html).  Sklearn supports parallel processing using the n_jobs parameter.  

## Explanation of Files
- `Data.csv` - Raw data, requires manual extraction to a `Data` folder.
- `/Logs/` - Contains logs for Go, R, and Python codes, including data preparation logs.
- `Data_Prep.ipynb` - Jupyter notebook for data preparation.
- `RF.go`, `RF_test.go`, `RF.r`, `RF.py` - Random forest implemenation files in Golang, R, and Python.
- `RF_go.exe` - Executable for the Golang implementation.

## Areas of Further Research / Enhancement
The underlying data in this analysis is imbalanced, with only 13% of the original corpus being spam.  As the Golang library being used does not support class weighting, no effort was made to address the data imbalance in this exercise.  A more accurate prediction could be achieved by applying an under sampling approach before model training. 

The golang random forest library used in this exercise has limited ability to adjust parameterization compared to similar random forest libraries in Python and R.  Many users involved in anything other than preliminary/light modeling will want significantly more flexibility to hyper tune parameters.  As such, users may want to explore other random forest libraries in Golang that allow for greater adjustment to underlying parameters or develop their own random forest application in Go.  

The malaschitz Random Forest library is another random forest go library: https://github.com/malaschitz/randomForest.  This library should be tested to see if this library would perform better than the Andy Sun’s library, or if additional parameters are available.  

The Andy Sun random forest library is not well documented.  My interpretation of the parametrization (based on a review of the underlying code library) may differ from reality and distort the findings and results from above.
