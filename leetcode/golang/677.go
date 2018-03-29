/*
677. Map Sum Pairs
Implement a MapSum class with insert, and sum methods.

For the method insert, you'll be given a pair of (string, integer). The string represents the key and the integer represents the value. If the key already existed, then the original key-value pair will be overridden to the new one.

For the method sum, you'll be given a string representing the prefix, and you need to return the sum of all the pairs' value whose key starts with the prefix.
*/

type MapSum struct {
    Mapping map[string]int 
};

/** Initialize your data structure here. */
func Constructor() MapSum {
    mapping := map[string]int{}

    return MapSum{Mapping:mapping}
}

func (this *MapSum) Insert(key string, val int)  {
    this.Mapping[key] = val 
}

func (this *MapSum) Sum(prefix string) int {
    prefixRegexp := regexp.MustCompile("^" + prefix)
    counter := 0
    
    for key, val := range(this.Mapping) {
        if prefixRegexp.Match[]byte(key)) {
           counter += val 
        } 
    }
    return counter
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
