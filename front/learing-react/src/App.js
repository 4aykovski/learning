import React, {useState} from 'react';
import './styles/App.css'
import PostList from "./components/PostList";


function App() {
    const [posts, setPosts] = useState([
        {id: 1, title: "Javascript", body: "Description"},
        {id: 2, title: "Golang", body: "Description"},
        {id: 3, title: "Python", body: "Description"},
    ])


    return (
        <div className="App">
            <form>
                <input type="text" placeholder="Post title"/>
                <input type="text" placeholder="Post description"/>
                <button>Create post</button>
            </form>
            <PostList posts={posts} title="Posts list"/>
        </div>
    );
}

export default App;
