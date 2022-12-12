#include <iostream>
#include <vector>
#include <math.h>

int first_part();

int second_part();

std::vector<char> get_moves();

bool in_path(std::vector<std::vector<int>> path, std::vector<int> vec);

std::vector<int> get_direction(char dir);

std::vector<int> add_vectors(std::vector<int> vec1, std::vector<int> vec2);

std::vector<int> get_shortest(std::vector<int> head, std::vector<int> tail);

double vector_distance(std::vector<int> vec1, std::vector<int> vec2);

int main()
{
    //std::cout << "First part: " << first_part() << std::endl;
    std::cout << "Second part: " << second_part() << std::endl;

    return 0;
}

int first_part()
{
    std::vector<char> moves = get_moves();
    std::vector<std::vector<int>> path;

    std::vector<int> head = { 0, 0 }, tail = { 0, 0 };
    path.push_back(tail);
    
    for (auto &move : moves)
    {
        head = add_vectors(head, get_direction(move));

        if (vector_distance(head, tail) < 2)
            continue;

        tail = get_shortest(head, tail);

        if (in_path(path, tail))
            continue;

        path.push_back(tail);
    }
    
    return path.size();
}

int second_part()
{
    std::vector<char> moves = get_moves();
    std::vector<std::vector<int>> tailHistory;

    std::vector<int> head = { 0, 0 }, tail = { 0, 0 };
    tailHistory.push_back(tail);
    
    for (auto &move : moves)
    {
        head = add_vectors(head, get_direction(move));

        if (vector_distance(head, tail) < 2)
            continue;

        tail = get_shortest(head, tail);

        tailHistory.push_back(tail);
    }
    
    int count = 0;
    bool contained;

    for (int i = 0; i < tailHistory.size() - 10; ++i)
    {
        contained = false;
        for (int j = 0; j < i; ++j)
            if (tailHistory.at(i) == tailHistory.at(j))
                contained = true;
        count += contained;
    }

    return count;
}

std::vector<char> get_moves()
{
    std::vector<char> moves;
    char move;
    int count;

    while (std::cin.peek() != EOF)
    {
        std::cin >> move >> count;
        for (; count; --count)
            moves.push_back(move);
    }

    return moves;
}

bool in_path(std::vector<std::vector<int>> path, std::vector<int> vec)
{
    for (auto &p : path)
        if (p == vec)
            return true;
    return false;
}

std::vector<int> get_direction(char dir)
{
    switch (dir)
    {
        case 'R':
            return {1, 0};
        case 'L':
            return {-1, 0};
        case 'U':
            return {0, -1};
        case 'D':
            return {0, 1};
        default:
            return {0, 0};
    }
}

std::vector<int> add_vectors(std::vector<int> vec1, std::vector<int> vec2)
{
    return { vec1[0] + vec2[0], vec1[1] + vec2[1] };
}

std::vector<int> get_shortest(std::vector<int> head, std::vector<int> tail)
{
    if (head == tail)
        return { 0, 0 };
    
    std::vector<std::vector<int>> moves;

    for (int i = -1; i <= 1; ++i)
        for (int j = -1; j <= 1; ++j)
            moves.push_back(add_vectors(tail, { i, j }));

    int min = 0;
    for (int i = 1; i < moves.size(); ++i)
        if (vector_distance(head, moves[i]) < vector_distance(head, moves[min]))
            min = i;

    return moves[min];
}

double vector_distance(std::vector<int> vec1, std::vector<int> vec2)
{
    return sqrt(pow(vec1[0] - vec2[0], 2) + pow(vec1[1] - vec2[1], 2));
}
