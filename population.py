from individual import Individual
import evaluation


class Population:

    def __init__(self, cipher, size, size_individual):
        self.individuals = []
        self.score = 0.0
        self.cipher = cipher
        for i in range(size):
            self.individuals.append(Individual(size_individual))
                
    def __repr__(self):
        return 'Best: {} Average score: {}'.format(self.best(), self.score/len(self.individuals))

    def sort(self):
        self.individuals.sort(key=lambda ind: ind.score)
        self.individuals.reverse()

    def evaluation(self):
        self.score = 0.0
        for i in self.individuals:
            i.score = evaluation.score(evaluation.permute(self.cipher, i.sequence))
            self.score += i.score
        
    def best(self):
        return self.individuals[0]

    def plaintext(self):
        return evaluation.permute(self.cipher, self.best().sequence)

    def display(self):
        for i in self.individuals:
            print i

    def crossing(self):
        half = int(len(self.individuals)/2)
        k = 0
        for i in range(half, len(self.individuals)):
            self.individuals[i] = self.individuals[k].crossing(self.individuals[k+1])
            k += 1

    def generation(self):
        self.crossing()
        self.evaluation()
        self.sort()
