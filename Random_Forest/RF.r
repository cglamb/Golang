#!/usr/bin/env Rscript
# loads in cleansed data and performs a random forest regression

#installs libraries if not already installed
if (!require("readr")) install.packages("readr", repos = "http://cran.us.r-project.org")  
if (!require("randomForest")) install.packages("randomForest", repos = "http://cran.us.r-project.org")
if (!require("caret")) install.packages("caret", repos = "http://cran.us.r-project.org")

# load libraries
library(readr)  #read in csv files
library(randomForest) #Breiman and Cutler's random forests for classification and regression: 	https://www.stat.berkeley.edu/~breiman/RandomForests/
library(caret)  #computes the confusion matrix

#initialize memory usage and timestamp
gc()
start.time <- proc.time()

# read in test and train data
x_test <- read_csv("Data/x_test.csv")
x_train <- read_csv("Data/x_train.csv")
y_test <- read_csv("Data/y_test.csv")
y_train <- read_csv("Data/y_train.csv")
y_train$y <- as.factor(y_train$y) #convert to factor for randomForest
y_test$y <- as.factor(y_test$y) #convert to factor for randomForest


# apply the random forest regression
rf2 <- randomForest(x = x_train, y = y_train$y,
                    ntree = 10,
                    mtry = 10
                    #ntree = 1577,
                    #mtry = sqrt(ncol(x_train)),
                    #nodesize = 1,
                    #min.node.size = 10,
                    #max.depth = 307,
                    #replace = FALSE,
                    #importance = TRUE,
                    #classwt = "balanced",
                    #randomState = 9)
)

# # making predictions on test data
rft2 <- predict(rf2, x_test)
rft2 <- factor(rft2, levels = levels(y_test$y)) 

# bundling accuracy, precision, and recall into a function
class_metrics <- function(method, actual) {
  confusion <- confusionMatrix(actual, method)
  accuracy <- confusion$overall['Accuracy']
  precision <- confusion$byClass['Pos Pred Value']
  recall <- confusion$byClass['Sensitivity']
  f1 <- 2 * ((precision * recall) / (precision + recall))
  
  cat("Accuracy:", accuracy, "\n")
  cat("Precision:", precision, "\n")
  cat("Recall:", recall, "\n")
  cat("F1:", f1, "\n")
}

#evaluating fit on the test data
print("applying class_metrics function")
class_metrics(rft2, y_test$y)

#end time and memory usage
end.time <- proc.time()


#print time and memory usage
print("Time and memory usage:")
print(end.time - start.time)
