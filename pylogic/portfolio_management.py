import pandas_datareader.data as web
import datetime as dt
import pandas as pd
import scipy.optimize as sc
import numpy as np
import plotly.graph_objects as go

# Import data
def GetData(stocks, start, end):
    print(f"Fetching data for stocks: {stocks}")
    stockData = web.DataReader(stocks, 'yahoo', start, end)
    return stockData

# Calculates the expected return and standard deviation (volatility) of the portfolio based on the given weights.
def portfolioPerformance(weights, meanReturns, covMatrix):
    returns = np.sum(meanReturns * weights) * 252
    std = np.sqrt(np.dot(weights.T, np.dot(covMatrix, weights))) * np.sqrt(252)
    return returns, std

# Returns the negative of the Sharpe Ratio (which should be maximized) for the portfolio.
def negativeSr(weights, meanReturns, covMatrix, riskFreeRate=0):
    pReturns, pStd = portfolioPerformance(weights, meanReturns, covMatrix)
    return - (pReturns - riskFreeRate) / pStd

# Uses optimization to find the portfolio weights that maximize the Sharpe Ratio.
def MaxSR(meanReturns, covMatrix, riskFreeRate=0, constraintSet=(0,1)):
    numAssets = len(meanReturns)
    args = (meanReturns, covMatrix, riskFreeRate)
    constraints = ({'type': 'eq', 'fun': lambda x: np.sum(x) - 1})
    bounds = tuple(constraintSet for _ in range(numAssets))
    result = sc.minimize(negativeSr, numAssets * [1. / numAssets], args=args, method='SLSQP', bounds=bounds, constraints=constraints)
    return result

def portfolioVariance(weights, meanReturns, covMatrix):
    return portfolioPerformance(weights, meanReturns, covMatrix)[1]

def minimizeVariance(meanReturns, covMatrix, constraintSet=(0,1)):
    numAssets = len(meanReturns)
    args = (meanReturns, covMatrix)
    constraints = ({'type': 'eq', 'fun': lambda x: np.sum(x) - 1})
    bounds = tuple(constraintSet for _ in range(numAssets))
    result = sc.minimize(portfolioVariance, numAssets * [1. / numAssets], args=args, method='SLSQP', bounds=bounds, constraints=constraints)
    return result

stockList = ['Nasdaq', 'SPY', 'US 20yr Bond']
# Efficient Frontier
stocks = [stock + '.AX' for stock in stockList]
endDate = dt.datetime.now()
startDate = endDate - dt.timedelta(days=365)

data = GetData(stocks, startDate, endDate)
meanReturns = data.mean()
covMatrix = data.cov()

def portfolioReturn(weights, meanReturns, covMatrix):
    return portfolioPerformance(weights, meanReturns, covMatrix)[0]

def efficientOpt(meanReturns, covMatrix, returnTarget, constraintSet=(0,1)):
    numAssets = len(meanReturns)
    args = (meanReturns, covMatrix)
    constraints = ({'type': 'eq', 'fun': lambda x: portfolioReturn(x, meanReturns, covMatrix) - returnTarget},
                   {'type': 'eq', 'fun': lambda x: np.sum(x) - 1})
    bounds = tuple(constraintSet for _ in range(numAssets))
    result = sc.minimize(portfolioVariance, numAssets * [1. / numAssets], args=args, method='SLSQP', bounds=bounds, constraints=constraints)
    return result

def calculatedResults(meanReturns, covMatrix, riskFreeRate=0, constraintSet=(0,1)):
    maxSR_Portfolio = MaxSR(meanReturns, covMatrix)
    maxSR_returns, maxSR_std = portfolioPerformance(maxSR_Portfolio.x, meanReturns, covMatrix)
    minVol_Portfolio = minimizeVariance(meanReturns, covMatrix)
    minVol_returns, minVol_std = portfolioPerformance(minVol_Portfolio.x, meanReturns, covMatrix)
    efficientList = []
    targetReturns = np.linspace(minVol_returns, maxSR_returns, 20)
    for target in targetReturns:
        efficientList.append(efficientOpt(meanReturns, covMatrix, target).fun)
    return maxSR_returns, maxSR_std, minVol_returns, minVol_std, efficientList

def Ef_graph(meanReturns, covMatrix, riskFreeRate=0, constraintSet=(0,1)):
    maxSR_returns, maxSR_std, minVol_returns, minVol_std, efficientList = calculatedResults(meanReturns, covMatrix)
    
    maxSharpeRatio = go.Scatter(
        name='Maximum Sharpe Ratio',
        mode='markers',
        x=[maxSR_std],
        y=[maxSR_returns],
        marker=dict(color='red', size=14, line=dict(width=3, color='black'))
    )
    
    MinVol = go.Scatter(
        name='Minimum Volatility',
        mode='markers',
        x=[minVol_std],
        y=[minVol_returns],
        marker=dict(color='green', size=14, line=dict(width=3, color='black'))
    )
    
    Ef_curve = go.Scatter(
        name="Efficient Frontier",
        mode='lines',
        x=[round(ef_std * 100, 2) for ef_std in efficientList],
        y=[round(ef_return * 100, 2) for ef_return in efficientList],
        line=dict(color='black', width=4, dash='dashdot')
    )
    
    data = [maxSharpeRatio, MinVol, Ef_curve]
    layout = go.Layout(
        title='Portfolio Optimization with the Efficient Frontier',
        yaxis=dict(title='Annualized Return (%)'),
        xaxis=dict(title='Annualized Volatility (%)'),
        showlegend=True
    )
    
    fig = go.Figure(data=data, layout=layout)
    return fig.show()
