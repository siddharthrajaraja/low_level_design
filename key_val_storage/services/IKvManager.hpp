#pragma once
#include <string>

using namespace std;

class IKVManager {
public:
    virtual void put(const string& key, const string& value) = 0;
    virtual string get(const string& key) = 0;
    virtual void remove(const string& key) = 0;
    virtual ~IKVManager() = default;
};