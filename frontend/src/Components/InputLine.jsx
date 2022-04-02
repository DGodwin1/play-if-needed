import React from "react";
import axios from "axios";

const InputLine = () => {
    // A component with a player input and an 'add'
    // but so that someone can add many users.
    let david = {
        ["name"]: "david",
        ["rank"]: 1,
    };

    let kevin = {
        ["name"]: "kevin",
        ["rank"]: 2,
    };

    let players = [david, kevin];

    return (
        <>
            <input type="text" placeholder="Players" />
            <button
                onClick={() => {
                    console.log("clicked");
                    console.log(JSON.stringify(david));
                    axios
                        .post(
                            "http://localhost:9001/sort-pairings",
                            JSON.stringify(players)
                        )
                        .then((r) => {
                            console.log(r.data);
                            console.log(r.status);
                        });
                }}>
                click
            </button>
        </>
    );
};

export default InputLine;
