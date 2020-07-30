#! /usr/bin/env python3

import matplotlib.pyplot as plt
from matplotlib.ticker import ScalarFormatter
import csv

def main():
    data = []
    time = []
    with open('tmp.csv') as csvfile:
        csvreader = csv.reader(csvfile)    
        for row in csvreader:
            time.append(int(row[0]))
            data.append(int(row[1]))

    print(data[len(data)-1])
    y_formatter = ScalarFormatter(useOffset=False)
    plt.axes().yaxis.set_major_formatter(y_formatter)
    plt.plot(time, data)
    plt.ylim(min(data), max(data))
    
    plt.show()

main()