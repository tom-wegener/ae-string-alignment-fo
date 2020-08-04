#! /usr/bin/env python3

import matplotlib.pyplot as plt
from matplotlib.ticker import ScalarFormatter
import csv


def main():
    data0 = []
    data1 = []
    data2 = []
    data3 = []
    data4 = []
    data5 = []
    time = []
    fileName = "80pop_20000gen_2geg_0.05druck_tpc"
    with open(fileName + ".csv") as csvfile:
        csvreader = csv.reader(csvfile)
        for row in csvreader:
            time.append(int(row[0]))
            data0.append(int(row[1]))

    fileName = "80pop_20000gen_3geg_0.05druck_tpc"
    with open(fileName + ".csv") as csvfile:
        csvreader = csv.reader(csvfile)
        for row in csvreader:
            data1.append(int(row[1]))

    fileName = "80pop_20000gen_2geg_0.05druck_opc_1"
    with open(fileName + ".csv") as csvfile:
        csvreader = csv.reader(csvfile)
        for row in csvreader:
            data2.append(int(row[1]))

    fileName = "80pop_20000gen_3geg_0.05druck_opc"
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

    plt.plot(time, data2, 'g', label='one Point Crossover, 2 Gegner')
    plt.plot(time, data3, 'y', label='one Point Crossover, 3 Gegner')
    plt.plot(time, data0, 'r', label='two Point Crossover, 2 Gegner')
    plt.plot(time, data1, 'b', label='two Point Crossover, 3 Gegner')
    plt.xlabel("Generation")
    plt.ylabel("Güte-Wert")
    maxAll = [max(data2), max(data1), max(data2), max(data3)]
    plt.ylim(0, max(maxAll))
    plt.xlim(0, max(time))
    plt.legend(loc="upper right")

    plt.savefig("opc_tpc.png", bbox_inches="tight")


main()
