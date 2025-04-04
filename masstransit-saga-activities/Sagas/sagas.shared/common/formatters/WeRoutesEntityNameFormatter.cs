using System.Reflection;
using MassTransit;

namespace weroutes.infrastructure.common.masstransitformatters;

public class WeRoutesEntityNameFormatter : IEntityNameFormatter
{
    private readonly MethodInfo formatMethod;
    private readonly IEntityNameFormatter _formatter;

    public WeRoutesEntityNameFormatter(IEntityNameFormatter formatter)
    {
        formatMethod = typeof(WeRoutesEntityNameFormatter).GetMethod(nameof(FormatEntityName))
            ?? throw new Exception($"Could not get method info of {nameof(FormatEntityName)}");
        _formatter = formatter;
    }

    public string FormatEntityName<T>()
    {
        Type type = typeof(T);


        if (type.IsAssignableTo(typeof(Fault)))
        {
            if (type.IsGenericType)
            {
                List<string> names = new() { "wr:fault" };

                foreach (var genericType in type.GetGenericArguments())
                {
                    names.Add(
                        (string)formatMethod.MakeGenericMethod(genericType).Invoke(this, Array.Empty<object>())!
                        );
                }

                return string.Join(":", names);
            }
            else
            {
                return _formatter.FormatEntityName<T>();
            }
        }

        ChannelNameAttribute channelNameAttr = 
            type.GetCustomAttribute<ChannelNameAttribute>()
            ?? throw new Exception($"The type {type.FullName} is not decorated with {typeof(ChannelNameAttribute).FullName}");

        string name = channelNameAttr.Name;

        return $"wr:mt:{name}";
    }
}
