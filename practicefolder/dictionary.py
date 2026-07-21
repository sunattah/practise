def remove_duplicates(text):
    textsplit = text.split()
    counts = {}
    for x in textsplit:
        if textsplit not in counts:
            textsplit.append(x)
    return counts

print(remove_duplicates([1, 3, 4, 5, 5, 5, 5, 6, 6]))
