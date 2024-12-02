#!/usr/bin/env python
# coding: utf-8

# loads in cleansed data and performs a random forest regression

# load libraries
import pandas as pd
import numpy as np
import itertools
import collections
import re
import nltk
import string
import opendatasets as od
import pickle
import time
from nltk.corpus import stopwords
from nltk import bigrams
from nltk.stem.porter import PorterStemmer
import tensorflow as tf
from tensorflow import keras
from sklearn.model_selection import train_test_split, GridSearchCV
from sklearn import metrics, svm
from sklearn.metrics import precision_score, recall_score, roc_curve, confusion_matrix, jaccard_score, f1_score
from sklearn.linear_model import LogisticRegression
from sklearn.naive_bayes import GaussianNB, BernoulliNB, MultinomialNB
from keras.layers import SimpleRNN, LSTM, Dense, Dropout, Activation, Flatten
from sklearn.preprocessing import LabelEncoder, OneHotEncoder
from sklearn.feature_extraction.text import TfidfVectorizer, CountVectorizer
from sklearn.ensemble import RandomForestClassifier, GradientBoostingClassifier, ExtraTreesClassifier, AdaBoostClassifier, AdaBoostClassifier
from xgboost import XGBClassifier, XGBRFClassifier
from sklearn.model_selection import RandomizedSearchCV
from imblearn.under_sampling import RandomUnderSampler
from imblearn.pipeline import Pipeline
from hyperopt import STATUS_OK, Trials, fmin, hp, tpe


#start timer
start_time = time.time()

#read in data
x_test = pd.read_csv('Data/x_test.csv')
x_train = pd.read_csv('Data/x_train.csv')
y_test = pd.read_csv('Data/y_test.csv')
y_train = pd.read_csv('Data/y_train.csv')
y_test = y_test.to_numpy().ravel()
y_train = y_train.to_numpy().ravel()


#fitting model using hypertuned parameters
rf2 = RandomForestClassifier(
    # class_weight="balanced",
    # random_state=9,
    n_estimators=1000,
    max_features=100,
    n_jobs=-1,
    warm_start=False
    # min_samples_split=10,
    # min_samples_leaf=1,
    # max_features='sqrt',
    # max_depth=307,
    # bootstrap=False)
)
rfm2 = rf2.fit(x_train, y_train)
rfr2 = rfm2.predict(x_train)
rft2 = rfm2.predict(x_test)

#passing function that calculates accuracy, preccsion, and recall on predictions made against test
def class_metrics(method):
    a = print("Accuracy:",metrics.accuracy_score(y_test, method))
    p = print("Precision:",metrics.precision_score(y_test, method))
    r =print("Recall:",metrics.recall_score(y_test, method))
    f =print("F1:",2*(metrics.precision_score(y_test, method)*metrics.recall_score(y_test, method))/(metrics.precision_score(y_test, method)+metrics.recall_score(y_test, method)))
    return a, p, r, f;


class_metrics(rft2)

#looking at default parameters
print(rf2.get_params())
print('Max Depth of any Tree: ',max([estimator.tree_.max_depth for estimator in rf2.estimators_]))

# Calculate and print the execution time
end_time = time.time()
execution_time = end_time - start_time
print(f"Execution Time: {execution_time} seconds")