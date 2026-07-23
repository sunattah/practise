# def remove_duplicates(text):
#     textsplit = text.split()
#     counts = {}
#     for x in textsplit:
#         if textsplit not in counts:
#             textsplit.append(x)
#         return counts
# print(remove_duplicates([1, 3, 4, 5, 5, 5, 5, 6, 6]))

# letter = ["a", "a", "a", "z", "z", "z", "b", "b", "b", "c"]
# count = {}
# for i in letter:
#     count[i] = count.get(i , 0) + 1
# print(count)


sentence = "hello world"
count = {}
for letter in sentence:
    count[letter] = count.get(letter, 0) + 1
print(count)