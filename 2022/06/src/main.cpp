#include <iostream>
#include <vector>

/// @brief Gets key that doesn't contain repeated chars
/// @param n key length
/// @return how many chars it had to read
int get_key(int n);

/// @brief Checks if string contains repeated char
/// @param key input string
/// @return true if contains, else false
bool has_repeats(std::string key);

int main()
{
    //std::cout << "First part: " << get_key(4) << std::endl;
    std::cout << "Second part: " << get_key(14) << std::endl;
}

int get_key(int n)
{
    std::string line;
    getline(std::cin, line);

    std::string key = "";

    for (int i = 0; i < line.length(); ++i)
    {
        key += line[i];
        if (key.length() < n)
            continue;
        if (!has_repeats(key))
            return i + 1;
        key.erase(0, 1);
    }

    return -1;
}

bool has_repeats(std::string key)
{
    for (int i = 0; i < key.length(); ++i)
        for (int j = 0; j < key.length(); ++j)
            if (i != j && key[i] == key[j])
                return true;
    return false;
}
