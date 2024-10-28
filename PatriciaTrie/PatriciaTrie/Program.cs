// See https://aka.ms/new-console-template for more information


using System.Diagnostics;
using System.Text.Json;

class Program
{
    static void Main(string[] args)
    {
        List<string> source = Enumerable.Range(0, 1000_000)
            .Select(x => Guid.NewGuid().ToString()).ToList();
        
        PatriciaTrie trie = new PatriciaTrie(new ConsolePrinter());

        foreach (var word in source)
        {
            trie.Insert(word);
        }

        // trie.Print();

        Random rnd = new Random(1035214);
        int rnd1 = rnd.Next(10000, 90000);
        int rnd2 = rnd.Next(10000, 90000);
        int rnd3 = rnd.Next(10000, 90000);
        int rnd4 = rnd.Next(10000, 90000);
        
        var toFind1 = source[rnd1];
        var toFind2 = source[rnd2];
        var toFind3 = source[rnd3];
        var toFind4 = source[rnd4];
        Console.WriteLine(rnd1);
        Console.WriteLine(rnd2);
        Console.WriteLine(rnd3);
        Console.WriteLine(rnd4);
        Stopwatch sw = new();
        sw.Start();
        for (int i = 0; i < 100; i++)
        {
            trie.Search(toFind1);
            trie.Search(toFind2);
            trie.Search(toFind3);
            trie.Search(toFind4);
        }
        sw.Stop();
        Console.WriteLine(sw.Elapsed.TotalMilliseconds);
        
        sw.Restart();
        for (int i = 0; i < 100; i++)
        {
            source.Any(x => x == toFind1);
            source.Any(x => x == toFind2);
            source.Any(x => x == toFind3);
            source.Any(x => x == toFind4);
        }
        sw.Stop();
        Console.WriteLine(sw.Elapsed.TotalMilliseconds);
        
        sw.Restart();
        for (int i = 0; i < 100; i++)
        {
            source.Contains(toFind1);
            source.Contains(toFind2);
            source.Contains(toFind3);
            source.Contains(toFind4);
        }
        sw.Stop();
        Console.WriteLine(sw.Elapsed.TotalMilliseconds);
        // Search for some words
        Console.WriteLine($"Search 'hello': {trie.Search(toFind1)}");
        Console.WriteLine($"Search 'hero': {trie.Search("hero")}");
        Console.WriteLine($"Search 'hey': {trie.Search("hey")}");
    }
}


public class PatriciaTrieNode
{
    public string Key { get; set; }
    public SortedDictionary<string, PatriciaTrieNode> Children { get; set; }

    public PatriciaTrieNode(string key)
    {
        Key = key;
        Children = new SortedDictionary<string, PatriciaTrieNode>(
            StringComparer.CurrentCultureIgnoreCase);
    }

    public bool IsLeaf => Children.Count == 0;
}

public abstract class Printer
{
    public abstract void Print(string value);
}

public class NoCollect : Printer
{
    public override void Print(string value)
    {
        
    }
}
public class CollectPrinter : Printer
{
    public List<string> Nodes { get; set; } = [];
    public override void Print(string value)
    {
        Nodes.Add(value);
    }

    public string GetString()
    {
        //We could just join, but I want to compare size of similar structures
        return JsonSerializer.Serialize(Nodes);
    }
}

public class ConsolePrinter : Printer
{
    public override void Print(string value)
    {
        Console.WriteLine(value);
    }
}

public class PatriciaTrie
{
    private readonly Printer _printer;
    private PatriciaTrieNode root;

    public PatriciaTrie(Printer printer)
    {
        _printer = printer;
        root = new PatriciaTrieNode(string.Empty);
    }

    // To provide a better grouping of random data
    static string SumToTwoDigits(string s)
    {
        int size = 3;
        // Step 1: Convert each character to its ASCII value and sum them
        int totalSum = s.Sum(c => (int)c);
        var str = totalSum.ToString().PadLeft(size, '0');
        return str[..size] + "@" + s;
    }
    
    public void Insert(string word)
    {
        word = SumToTwoDigits(word);
        
        var currentNode = root;
        while (!string.IsNullOrEmpty(word))
        {
            bool prefixFound = false;
            foreach (var child in currentNode.Children)
            {
                var commonPrefix = GetCommonPrefix(word, child.Key);
                if (!string.IsNullOrEmpty(commonPrefix))
                {
                    prefixFound = true;
                    if (commonPrefix == child.Key)
                    {
                        currentNode = child.Value;
                        word = word.Substring(commonPrefix.Length);
                    }
                    else
                    {
                        var remainingKey = child.Key.Substring(commonPrefix.Length);
                        var newNode = new PatriciaTrieNode(commonPrefix);
                        newNode.Children[remainingKey] = child.Value;

                        currentNode.Children.Remove(child.Key);
                        currentNode.Children[commonPrefix] = newNode;

                        if (!string.IsNullOrEmpty(word.Substring(commonPrefix.Length)))
                        {
                            newNode.Children[word.Substring(commonPrefix.Length)] = new PatriciaTrieNode(word);
                        }
                        return;
                    }
                    break;
                }
            }

            if (!prefixFound)
            {
                currentNode.Children[word] = new PatriciaTrieNode(word);
                return;
            }
        }
    }

    // Search for a word in the Patricia trie
    public bool Search(string word)
    {
        word = SumToTwoDigits(word);
        var currentNode = root;
        bool found = false;
        while (!string.IsNullOrEmpty(word))
        {
            found = false;
            foreach (KeyValuePair<string, PatriciaTrieNode> child in currentNode.Children)
            {
                if (WordStartsWith(word, child.Key))
                {
                    currentNode = child.Value;
                    word = word[child.Key.Length..];
                    found = true;
                    break;
                }
            }
            if (!found) return false;
        }
        return currentNode.IsLeaf;
    }

    private static bool WordStartsWith(ReadOnlySpan<char> first, ReadOnlySpan<char> key)
    {
        if (first.Length < key.Length)
        {
            return false;
        }

        for (int i = 0; i < key.Length; i++)
        {
            if (first[i] != key[i])
            {
                return false;
            }
        }

        return true;
    }

    // Helper function to get the common prefix between two strings
    private string GetCommonPrefix(string str1, string str2)
    {
        int minLength = Math.Min(str1.Length, str2.Length);
        int i = 0;
        while (i < minLength && str1[i] == str2[i])
        {
            i++;
        }
        return str1.Substring(0, i);
    }

    // Print Patricia Trie for debugging
    public void Print()
    {
        PrintNode(root, 0);
    }

    private void PrintNode(PatriciaTrieNode node, int level)
    {
        foreach (var child in node.Children)
        {
            _printer.Print(new string('-', level) + child.Key);
            PrintNode(child.Value, level + 1);
        }
    }
}
