import React, {Component, useState} from 'react';
import {Text, View, StyleSheet, Button, TextInput, Image} from 'react-native';

// const Cat = () => {
//     return <Text>Hello, I am your cat!</Text>;
// };
//
// export default Cat;

// export function HelloWorld(props: GreetingsProps){
//     const [count, setCount] = useState(0);
//
//     return (
//       <View
//         style={styles.view}>
//           <Text>Hello {props.name}. You clicked {count} times</Text>
//           <Button onPress={() => setCount(count + 1)} title="Click Me"></Button>
//       </View>
//     );
// }


export class HelloWorld extends Component<GreetingsProps> {
    state = {
        count: 3,
    };

    onPress = () => {
        this.setState({
            count: this.state.count + 1,
        });
    };

    render() {
        return (
            <View
                style={styles.view}>
                <Text>Hello {getFullName(this.props)}. You clicked {this.state.count} times</Text>
                <Button onPress={this.onPress} title="Click Me"></Button>
                <Image source={sources.image} style={styles.image}/>
                <TextInput style={styles.text_input} defaultValue="type"></TextInput>
            </View>
        );
    }
}

type GreetingsProps = {
    first_name: string,
    last_name: string,
}

function getFullName(data: GreetingsProps): string {
    return `${data.first_name} ${data.last_name}`;
}

const sources = {
    image: {
        uri: 'https://reactnative.dev/docs/assets/p_cat1.png',
    }
}

const styles = StyleSheet.create({
    view: {
        flex: 1,
    },
    text_input: {
        height: 40,
        borderColor: 'gray',
        borderWidth: 1,
    },
    image: {
        width: 200,
        height: 200
    }
})


