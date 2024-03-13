import csv
import re


def nums(s):
    match = re.search(r'\d+', s)
    if match:
        return int(match.group())
    return float('inf')

def sort_csv(csv_file):
    with open(csv_file, 'r', newline='') as file:
        reader = csv.reader(file, delimiter=';')
        data = list(reader)

    sorted_data = sorted(data, key=lambda x: int(nums(x[0][6:])))

    with open(csv_file, 'w', newline='') as file:
        writer = csv.writer(file, delimiter=';')
        writer.writerows(sorted_data)

csv_file = "locations.csv"
sort_csv(csv_file)
print("CSV file sorted successfully.")