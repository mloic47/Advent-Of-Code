# read file
with open('Day01/listsDistanceInput.txt', 'r') as f:
    data = f.readlines()

leftList= []
rightList = []
for line in data:
    left,right = line.split()
    leftList.append(left)
    rightList.append(right)

listValues = []
# loop through the left list and count the values of the left found in the right list
for i in range(len(leftList)):
    counValue = rightList.count(leftList[i])
    multiply = int(leftList[i]) * counValue
    # append the values to a list
    listValues.append(multiply)

# sum the values in the list
sumValues = sum(listValues)
print(sumValues)