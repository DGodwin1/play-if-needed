import React from "react";

const InputLine = () => {
    // A component with a player input and an 'add'
    // but so that someone can add many users.
    return (
        <>
            <input type="text" placeholder="Players" />
            <input
                type="submit"
                method="post"
                value="Sort out the doubles"
                action="/someWebPage"
            />
        </>
    );
};

export default InputLine;
