import random

def generate_random_sequence(length):
    sequence = []
    while len(sequence) < length:
        index = random.randint(0, len(sequence))
        if index not in sequence:
            sequence.append(index)
    return sequence
