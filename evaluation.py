def permute(text, permu):
    permuted = []
    i = 0
    block = 0
    while i < len(text) - 13:
        permuted.append(text[block + permu[i % 13]])
        i += 1
        if i % 13 == 0:
            block += 13
    return ''.join(permuted)


def score(text):
    # BONUS
    bigrams = [('TH', 3.2), ('HE', 2.6), ('IN', 2.2), ('ER', 1.9), ('AN', 1.8), ('RE', 1.5), ('ES', 1.4), ('ON', 1.4), ('ST', 1.4), ('NT', 1.3), ('EN', 1.3), ('ED', 1.3), ('ND', 1.2), ('AT', 1.2), ('TI', 1.2), ('TE', 1.1), ('OR', 1.1), ('AR', 1), ('HA', 1), ('OF', 1)]
    trigrams = [('THE', 22.5), ('AND', 9.0), ('ING', 6.6), ('ENT', 5.4), ('ION', 4.5), ('NTH', 4.2), ('TER', 3.9), ('INT', 3.9), ('OFT', 3.9), ('THA', 3.9), ('ERE', 3.9), ('TIO', 3.6), ('HER', 3.6), ('FTH', 3.6), ('ETH',  3.3), ('ATI', 3.3), ('HAT', 3), ('ATE', 3), ('STH', 3), ('EST', 3)]
        
    # MALUS
    malus = [('TX', -1.0), ('TZ', -1.0), ('OZ', -1.0), ('IJ', -1.0), ('IY', -1.0)]
        
    elements = bigrams + trigrams + malus
        
    scored = 0.0
    for element, coefficient in elements:
        scored += coefficient * text.count(element)
    return scored
