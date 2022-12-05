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

/// @brief Second part of the puzzle
/// @return Sum of priorities of group badges
int second_part();

/// @brief Finds badge of given group
/// @param line group rucksacks items
/// @return priority of the badge
int process_group(std::string line[]);

/// @brief Checks if item is in rucksack
/// @param line rucksack items
/// @param c item
/// @return true if found, else false
bool in_line(std::string line, char c);

int main()
{
    //std::cout << "Sum of priorities: " << first_part() << std::endl;
    std::cout << "Sum of badges priorities: " << second_part() << std::endl;

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

int second_part()
{
    std::string line[3];
    int sum = 0;
    while (getline(std::cin, line[0]) &&
           getline(std::cin, line[1]) &&
           getline(std::cin, line[2]))
        sum += process_group(line);

    return sum;
}

int process_group(std::string line[])
{
    for (int i = 0; i < line[0].length(); ++i)
    {
        if (in_line(line[1], line[0][i]) &&
            in_line(line[2], line[0][i]))
            return get_priority(line[0][i]);
    }
    return 0;
}

bool in_line(std::string line, char c)
{
    for (int i = 0; i < line.length(); ++i)
        if (line[i] == c)
            return true;
    return false;
}
