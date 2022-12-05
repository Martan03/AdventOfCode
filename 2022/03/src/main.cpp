#include <iostream>

/// @brief First part of the puzzle
/// @return Sum of priorities of repeated items in each rucksack
int first_part();

/// @brief Gets repeated item priority value in given rucksack
/// @param line rucksack items
/// @return priority of repeated item
int get_repeated_item(std::string line);

/// @brief Gets priority of given item
/// @param c item
/// @return priority of item
int get_priority(char c);

int main()
{
    std::cout << "Sum of priorities: " << first_part() << std::endl;

    return 0;
}

int first_part()
{
    std::string line;
    int sum = 0;
    
    while (getline(std::cin, line))
        sum += get_repeated_item(line);

    return sum;
}

int get_repeated_item(std::string line)
{
    int lineMid = line.length() / 2;
    for (int i = 0; i < lineMid; ++i)
    {
        for (int j = lineMid; j < line.length(); ++j)
        {
            if (line[i] == line[j])
                return get_priority(line[j]);
        }
    }

    return 0;
}

int get_priority(char c)
{
    if (c >= 'A' && c <= 'Z')
        return c - 'A' + 27;
    return c - 'a' + 1;
}
