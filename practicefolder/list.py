# def remove_duplicates(items):
#     itemList = []
# for x in items:
#   if x  not in itemList:
#       itemList.append(x)
#       return itemList

# print(remove_duplicates([1,3,4,5,5,5,5,6,6,]))

def remove_duplicates(items):
    itemList = []
    for x in items:
        if x not in itemList:
            itemList.append(x)
    return itemList

print(remove_duplicates([1, 3, 4, 5, 5, 5, 5, 6, 6]))