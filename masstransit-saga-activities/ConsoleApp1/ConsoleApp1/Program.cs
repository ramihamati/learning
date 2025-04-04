// See https://aka.ms/new-console-template for more information


using System.Text.Json;

var p1 = new Person()
{
    Material =
    [
        new Material()
        {
            Name = "Blue"
        }
    ],
    Name = "White"
};

var p2 = p1 with { };
p2.Name = "p2Name";
p2.Material[0].Name = "p2Material";

Console.WriteLine(p1);
Console.WriteLine(p2);

record Material
{
    public string Name { get; set; }
}

record Person
{
    public string Name { get; set; }
    public Material[] Material { get; set; }

    public override string ToString()
    {
        return JsonSerializer.Serialize(this);
    }

    public Person(Person p)
    {
        Name = p.Name;
        Material = [
            ..p.Material.Select(x=> x with{}).ToArray()
        ];
    }
}