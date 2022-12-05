#include <iostream>

/// @brief First part of the puzzle
/// @return number of pairs that are contained in their pair
int first_part();

/// @brief Second part of the puzzle
/// @return number of pairs that overlap
int second_part();

int main()
{
    // std::cout << "First part: " << first_part() << std::endl;
    std::cout << "Second part: " << second_part() << std::endl;

    return 0;
}

int first_part()
{
    int xStart, xEnd, yStart, yEnd;
    char c;
    int count = 0;
    while (std::cin.peek() != EOF)
    {
        std::cin >> xStart >> c >> xEnd >> c >> yStart >> c >> yEnd;
        if ((xStart <= yStart && xEnd >= yEnd) ||
            (yStart <= xStart && yEnd >= xEnd))
            ++count;
    }

    return count;
}

int second_part()
{
    int xStart, xEnd, yStart, yEnd;
    char c;
    int count = 0;
    while (std::cin.peek() != EOF)
    {
        std::cin >> xStart >> c >> xEnd >> c >> yStart >> c >> yEnd;
        if ((xEnd >= yStart && xStart <= yStart) ||
            (yEnd >= xStart && yStart <= xStart))
            ++count;
    }

    return count;
}
 