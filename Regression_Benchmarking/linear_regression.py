#!/usr/bin/env python
# coding: utf-8

#query from thomas w miller

#importing libarries
from __future__ import division, print_function
import pandas as pd
import numpy as np
import statsmodels.api as sm
import matplotlib.pyplot as plt
from statsmodels.graphics.regressionplots import abline_plot
import time

#record clock time at start for benchmarking
start = time.time()

#bringing in data for modeling
anscombe = pd.DataFrame({'x1' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
'x2' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
'x3' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
'x4' : [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
'y1' : [8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26,10.84, 4.82, 5.68],
'y2' : [9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74],
'y3' : [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73],
'y4' : [6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89]})


# # Set 1 - Fitting line and outputing regression results
set_I_design_matrix = sm.add_constant(anscombe['x1'])
set_I_model = sm.OLS(anscombe['y1'], set_I_design_matrix)
print(set_I_model.fit().summary())

# # Set 2 - Fitting line and outputing regression results
set_II_design_matrix = sm.add_constant(anscombe['x2'])
set_II_model = sm.OLS(anscombe['y2'], set_II_design_matrix)
print(set_II_model.fit().summary())

# # Set 3 - Fitting line and outputing regression results
set_III_design_matrix = sm.add_constant(anscombe['x3'])
set_III_model = sm.OLS(anscombe['y3'], set_III_design_matrix)
print(set_III_model.fit().summary())


# # Set 4 - Fitting line and outputing regression results
set_IV_design_matrix = sm.add_constant(anscombe['x4'])
set_IV_model = sm.OLS(anscombe['y4'], set_IV_design_matrix)
print(set_IV_model.fit().summary())

#recording end time and output elapsed time
end = time.time()
print(end-start)