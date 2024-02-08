import React, {useRef, useState} from 'react';
import './styles/App.css'
import PostList from "./components/PostList";
import MyButton from "./components/UI/button/MyButton";
import MyInput from "./components/UI/input/MyInput";
import PostForm from "./components/postForm";


function App() {
    const [posts, setPosts] = useState([
        {id: 1, title: "Javascript", body: "Description"},
        {id: 2, title: "Golang", body: "Description"},
        {id: 3, title: "Python", body: "Description"},
    ])

    const createPost = (newPost) => {
        setPosts([...posts, newPost])
    }

    return (
        <div className="App">
            <PostForm create={createPost}/>
            <PostList
                posts={posts}
                title="Posts list"/>
        </div>
    );
}

export default App;
