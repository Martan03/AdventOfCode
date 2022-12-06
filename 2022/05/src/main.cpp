#include <iostream>
#include <vector>

std::string first_part();

std::string second_part();

int main()
{
    std::cout << second_part() << std::endl;

    return 0;
}

void print_vector(std::vector<std::vector<char>> vector)
{
    for (auto &column : vector)
    {
        for (auto &crate : column)
        {
            std::cout << crate;
        }
        std::cout << std::endl;
    }
}

std::string first_part()
{
    std::vector<std::string> lines;
    std::string line;

    while (getline(std::cin, line) && line != "")
        lines.push_back(line);
    lines.pop_back();

    int size = (lines.at(0).length() + 1) / 4;

    std::vector<std::vector<char>> crates(size);
    for (int i = lines.size() - 1; i >= 0; --i)
    {
        for (int j = 0; j < size; ++j)
        {
            if (lines[i][j * 4 + 1] != ' ')
                crates[j].push_back(lines[i][j * 4 + 1]);
        }
    }

    int count, from, to;
    while (std::cin.peek() != EOF)
    {
        std::cin >> line >> count >> line >> from >> line >> to;

        for (; count; --count)
        {
            crates[to - 1].push_back(crates[from - 1].back());
            crates[from - 1].pop_back();
        }
    }
    
    line = "";
    for (auto &column : crates)
        line += column.back();

    return line;
}

std::string second_part()
{
    std::vector<std::string> lines;
    std::string line;

    while (getline(std::cin, line) && line != "")
        lines.push_back(line);
    lines.pop_back();

    int size = (lines.at(0).length() + 1) / 4;

    std::vector<std::vector<char>> crates(size);
    for (int i = lines.size() - 1; i >= 0; --i)
    {
        for (int j = 0; j < size; ++j)
        {
            if (lines[i][j * 4 + 1] != ' ')
                crates[j].push_back(lines[i][j * 4 + 1]);
        }
    }

    int count, from, to;
    while (std::cin.peek() != EOF)
    {
        std::cin >> line >> count >> line >> from >> line >> to;

        for (int i = 0; i < count / 2; ++i)
        {
            char temp = crates[from - 1][crates[from - 1].size() - 1 - i];
            crates[from - 1][crates[from - 1].size() - 1 - i] = crates[from - 1][crates[from - 1].size() - count + i];
            crates[from - 1][crates[from - 1].size() - count + i] = temp;
        }
        
        for (; count; --count)
        {
            crates[to - 1].push_back(crates[from - 1].back());
            crates[from - 1].pop_back();
        }
    }
    
    line = "";
    for (auto &column : crates)
        line += column.back();

    return line;
}
