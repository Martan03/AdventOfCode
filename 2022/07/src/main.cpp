#include <iostream>
#include <vector>
#include <cstring>

int first_part();

int second_part();

std::vector<int> get_sizes();

int get_size(std::vector<std::vector<std::string>> tree, std::vector<std::string> folders, std::vector<int> size, int id);

int folder_index(std::vector<std::string> folders, std::string folder);

bool is_number(std::string str);

bool is_number(std::string str, int end);

int main()
{
    //std::cout << "First part: " << first_part() << std::endl;
    std::cout << "Second part: " << second_part() << std::endl;

    return 0;
}

int first_part()
{
    std::vector<int> size = get_sizes();

    int sum = 0;

    for (auto &s : size)
        if (s < 100000)
            sum += s;

    return sum;
}

int second_part()
{
    std::vector<int> size = get_sizes();

    int minSpace = 30000000 - (70000000 - size[0]);

    int min = -1;

    for (auto &s : size)
        if (s > minSpace && (s < min || min == -1))
            min = s;
    
    return min;
}

std::vector<int> get_sizes()
{
    std::string path = "";
    std::vector<std::string> folder;
    std::vector<int> size;
    std::vector<std::vector<std::string>> tree;
    std::string line;
    while (getline(std::cin, line))
    {
        if (line.find("dir", 0) == 0)
        {
            line.erase(0, line.find_first_of(' ') + 1);
            tree[folder.size() - 1].push_back(path + "/" + line);
            continue;
        }
        else if (is_number(line, line.find_first_of(' ')))
        {
            line.erase(line.find_first_of(' '), line.length());
            size[folder.size() - 1] += stoi(line);
            continue;
        }
        else if (line.find("$ cd", 0) != 0)
            continue;

        line.erase(0, line.find_last_of(' ') + 1);

        if (line == "..")
        {
            path.erase(path.find_last_of('/'), path.length());
            continue;
        }

        if (path != "")
            path += "/";

        path += line;
        
        folder.push_back(path);
        size.push_back(0);
        tree.push_back(std::vector<std::string>());
    }

    for (int i = 0; i < folder.size(); ++i)
        size[i] = get_size(tree, folder, size, i);

    return size;
}

int get_size(std::vector<std::vector<std::string>> tree, std::vector<std::string> folders, std::vector<int> size, int id)
{
    int s = 0;

    for (auto &item : tree[id])
    {
        int id = folder_index(folders, item);
        if (id == -1)
            continue;
        s += get_size(tree, folders, size, id);
    }
    return s + size[id];
}

int folder_index(std::vector<std::string> folders, std::string folder)
{
    for (int i = 0; i < folders.size(); ++i)
        if (folder == folders[i])
            return i;
    return -1;
}

bool is_number(std::string str)
{
    return is_number(str, str.length());
}

bool is_number(std::string str, int end)
{
    for (int i = 0; i < end && i < str.length(); ++i)
        if (!isdigit(str[i]))
            return false;
    return true;
}
