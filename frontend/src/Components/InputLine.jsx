import React from "react";
import axios from "axios";

const InputLine = () => {
    // A component with a player input and an 'add'
    // but so that someone can add many users.
    let david = {
        name: "david",
    };
    let kevin = {
        name: "kevin",
    };

    let players = [david, sophie];

    return (
        <>
            <input type="text" placeholder="Players" />
            <button
                onClick={() => {
                    console.log("clicked");
                    axios
                        .post("http://localhost:9001/sort-pairings", players)
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
