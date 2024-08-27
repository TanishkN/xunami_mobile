import yfinance as yf
import pandas as pd 
import numpy as np
import matplotlib.pyplot as plt
#End of Day Data.
stock = yf.download("AAPL", start = "2024-01-01", end = "2024-08-26")
print(stock)

us20yr = yf.download()
#xag

#Pre-Constructed Portfolios
