using MassTransit;
using System.Reflection;

namespace sagas.shared;

public class WeRoutesEndpointNameFormatter : IEndpointNameFormatter
{
    public string Separator { get; } = "-";
    private readonly Dictionary<Type, ChannelNameAttribute> _preregisteredAttributes;

    public WeRoutesEndpointNameFormatter()
    {
        _preregisteredAttributes = new();
    }

    public WeRoutesEndpointNameFormatter(
      Dictionary<Type, ChannelNameAttribute> preregisteredAttributes)
    {
        _preregisteredAttributes = preregisteredAttributes ?? new();
    }

    public string TemporaryEndpoint(string tag)
    {
        if (string.IsNullOrWhiteSpace(tag))
            tag = "endpoint";

        return $"wr:tmp:{tag}:{Guid.NewGuid()}";
    }

    public string Consumer<T>()
        where T : class, IConsumer
    {
        return FormatName("consumer", typeof(T));
    }

    public string Message<T>()
        where T : class
    {
        return FormatName("message", typeof(T));
    }

    public string Saga<T>()
        where T : class, ISaga
    {
        return FormatName("saga", typeof(T));
    }

    public string ExecuteActivity<T, TArguments>()
        where T : class, IExecuteActivity<TArguments>
        where TArguments : class
    {
        return FormatName("activity:execute", typeof(T));
    }

    public string CompensateActivity<T, TLog>()
        where T : class, ICompensateActivity<TLog>
        where TLog : class
    {
        return FormatName("activity:compensate", typeof(T));
    }

    public string SanitizeName(string name)
    {
        return name;
        //no longer sanitizing. 
        //if (name.Length> 0 && replacechars.Contains(name[0]))
        //{
        //    return SanitizeName(name[1..]);
        //}

        //return _pattern.Replace(name, m => Separator + m.Value).ToLowerInvariant();
    }

    private string FormatName(string prefix, Type type)
    {
        if (prefix.EndsWith(Separator))
        {
            prefix = prefix[0..^1];
        }

        ChannelNameAttribute? channelNameAttr =
            _preregisteredAttributes.ContainsKey(type)
            ? _preregisteredAttributes[type]
            : type.GetCustomAttribute<ChannelNameAttribute>();

        if (channelNameAttr is null)
        {
            throw new Exception($"The type {type.FullName} is not decorated with {typeof(ChannelNameAttribute).FullName}");
        }

        string name = channelNameAttr.Name;

        if (!name.StartsWith(prefix))
        {
            return $"wr:mt:{prefix}:{name}";
        }
        else
        {
            return $"wr:mt:{name}";
        }

        // no longer sanitizing because it's part of the user
        // to determine the naming standards
        //name = string.Join("", name.Select(x => replacechars.Contains(x) ? sepparator : x));

        // no longer envorcing the prefixx to have the format wr:mt:prefix:name
        // instead we can have wr:mt:prefix-name

        // and that's valid. i.e. wr:mt:consumer-domain:event-domain
        //    - in this example we will pas the string"consumer-domain:event-domain"
        //      and because "consumer" is the prefix in the old way this was detected and 
        //      changed to wr:mt:consumer:domain:Event-domain

        //if (name.StartsWith(prefix))
        //{
        //    name = name[prefix.Length..];
        //}

        //name = SanitizeName(name);

        //if (!name.StartsWith(prefix))
        //{
        //    return $"wr:mt:{prefix}:{name}";
        //}
        //else
        //{
        //    return $"wr:mt:{name}";
        //}
    }
}
