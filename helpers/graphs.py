#! /usr/bin/env python3

import matplotlib.pyplot as plt
from matplotlib.ticker import ScalarFormatter
import csv


def main():
    data = []
    time = []
    fileName = "2000pop_2000gen_5geg_0.05druck"
    with open(fileName + ".csv") as csvfile:
        csvreader = csv.reader(csvfile)
        for row in csvreader:
            time.append(int(row[0]))
            data.append(int(row[1]))

    fig_size = plt.rcParams["figure.figsize"]
    fig_size[0] = 14
    fig_size[1] = 8
    plt.rcParams["figure.figsize"] = fig_size

    y_formatter = ScalarFormatter(useOffset=False)
    plt.axes().yaxis.set_major_formatter(y_formatter)
    plt.plot(time, data)
    plt.xlabel("Generation")
    plt.ylabel("Güte-Wert")
    plt.ylim(0, max(data)+1000)
    plt.xlim(0, max(time))

    plt.savefig(fileName+".png", bbox_inches="tight")


main()
