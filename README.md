Transpogen
==========

Utilisation d'un algorithme génétique pour décrypter un chiffrement par transposition
-------------------------------------------------------------------------------------

#Algorithme génétique
* Un individu représente une solution possible
* Une population est un ensemble d'individus
* Chaque individu est noté selon la solution apportée
* A chaque génération les individus les moins performants sont remplacés par des croisements des individus les plus performants (avec possible mutation)

#Chiffrement par transposition
* Le texte chiffré est divisé par blocs (bloc de 13 caractères dans notre cas)
* Chaque bloc est un anagramme des caratères présents
* L'espace est comptabilisé comme caractère
* Il y a 13! permutations possibles

#Concrètement
* Chaque individu va représenter une permutation possible
* La note d'un individu se base sur le texte résultant de sa permutation
* La note est calculée selon les bigrammes/trigrammes présents dans le texte ainsi que des mots probables

Programmation en Python 2.7