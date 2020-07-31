#! /usr/bin/env python3

import matplotlib.pyplot as plt
from matplotlib.ticker import ScalarFormatter
import csv


def main():
    data0 = []
    data1 = []
    data2 = []
    data3 = []
    time = []
    fileName = "80pop_2000gen_2geg_0.05druck_opc"
    with open(fileName + ".csv") as csvfile:
        csvreader = csv.reader(csvfile)
        for row in csvreader:
            time.append(int(row[0]))
            data0.append(int(row[1]))

    fileName = "80pop_2000gen_3geg_0.05druck_opc"
    with open(fileName + ".csv") as csvfile:
        csvreader = csv.reader(csvfile)
        for row in csvreader:
            data1.append(int(row[1]))

    fileName = "80pop_2000gen_5geg_0.05druck_opc"
    with open(fileName + ".csv") as csvfile:
        csvreader = csv.reader(csvfile)
        for row in csvreader:
            data2.append(int(row[1]))

    fileName = "80pop_2000gen_7geg_0.05druck_opc"
    with open(fileName + ".csv") as csvfile:
        csvreader = csv.reader(csvfile)
        for row in csvreader:
            data3.append(int(row[1]))

    fig_size = plt.rcParams["figure.figsize"]
    fig_size[0] = 14
    fig_size[1] = 8
    plt.rcParams["figure.figsize"] = fig_size

    y_formatter = ScalarFormatter(useOffset=False)
    plt.axes().yaxis.set_major_formatter(y_formatter)
    


    plt.plot(time, data0, 'r', label='2 Gegner')
    plt.plot(time, data1, 'b',label='3 Gegner')
    plt.plot(time, data2, 'g',label='5 Gegner')
    plt.plot(time, data3,'y',label='7 Gegner')
    plt.xlabel("Generation")
    plt.ylabel("Güte-Wert")
    maxAll = [max(data1), max(data0)]
    plt.ylim(0, max(maxAll))
    plt.xlim(0, max(time))
    plt.legend(loc="upper right")

    plt.savefig("gegner.png", bbox_inches="tight")


main()
