import time
from population import Population


def transpogen(cipher, generations, individuals, size):
    start = time.time()
    print '[*] Genetic algorithm of {} individuals over {} generations'.format(individuals, generations)
    print '[*] Start at {}\n'.format(time.strftime('%X'))
    p = Population(cipher, individuals, size)

    for i in range(generations):
        p.generation()
        print '[+] ### Generation {} ### {}'.format(i, p)

    print '\n[*] End at {} ({} seconds)'.format(time.strftime('%X'), round(time.time() - start, 2))
    print '[*] Best solution is {}'.format(p.best())
    with open('plain.text', 'w') as solution_file:
        solution_file.write(p.plaintext())


if __name__ == '__main__':
    cipher = open('cipher.txt').read()
    transpogen(cipher, 50, 100, 13)
