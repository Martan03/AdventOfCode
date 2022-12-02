#include <iostream>

/// @brief Analyzes the games from stdin
/// @return final score
int Analyze();

/// @brief Gets score of the round
/// @param a first player
/// @param b second player
/// @return Round score
int GetScore(char a, char b);

/// @brief Gets score for playing specific figure
/// @param f figure character
/// @return figure score
int GetFigureScore(char f);

int main()
{
    std::cout << Analyze() << std::endl;

    return 0;
}

int Analyze()
{
    char a, b;
    int scoreSum = 0;
    while (std::cin.peek() != EOF)
    {
        std::cin >> a;
        std::cin >> b;
        scoreSum += GetScore(b, a);
    }
    
    return scoreSum;
}

int GetScore(char a, char b)
{
    int aVal = GetFigureScore(a), bVal = GetFigureScore(b);
    if (aVal == 1)
        return ((bVal - 1 > 0) ? bVal - 1: 3);
    else if (aVal == 3)
        return (bVal % 3) + 7;
    return bVal + 3;
}

int GetFigureScore(char f)
{
    switch (f)
    {
        // Rock
        case 'A':
        case 'X':
            return 1;
        // Paper
        case 'B':
        case 'Y':
            return 2;
        // Scissors
        case 'C':
        case 'Z':
            return 3;
        default:
            return 0;
    }
}