import random


class Individual:
    
    def __init__(self, size):
        self.sequence = []
        self.score = 0.0
        for i in range(size):
            self.sequence.append(i)
        random.shuffle(self.sequence)

    def __repr__(self):
        return '{} ({})'.format(self.sequence, self.score)
    
    def shift(self, n):
        tmp = []
        for i in range(len(self.sequence)):
            tmp.append(self.sequence[(i + n) % len(self.sequence)])
        self.sequence = tmp

    def mutation(self, p):
        if random.random() <= p:
            self.shift(random.randint(1, len(self.sequence)))

    def crossing(self, individual):
        cross = []
        if self.score >= individual.score:
            best = self
            weak = individual
        else:
            best = individual
            weak = self
        half = int(len(self.sequence)/2)
        for i in range(half):
            cross.append(best.sequence[i])
        for element in weak.sequence:
            if cross.count(element) != 1:
                cross.append(element)
        child = Individual(len(self.sequence))
        child.sequence = cross
        child.mutation(0.7)
        return child
