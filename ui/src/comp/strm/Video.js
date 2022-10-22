import React from 'react';

const items = ["flower", "ed", "buck"];

export default class Video extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            items: items,
        };
    }

    videoPath = (item) => {
        let path = "http://localhost:8003/" + item + ".mp4"
        // console.log(path)
        return path
    }

    render() {
        return (
            <div>
                <form className="pure-form searchVideoForm">
                    <input type="text" className="pure-input-1 searchVideoTextBox" placeholder="search video" onChange={evt => this.filterVideo(evt)} />
                    <br />
                </form>

                {this.state.items.map((item) =>
                    <div className="pure-u-1-1 pure-u-md-1-2 pure-u-lg-1-3 videoGrid" key={item}>
                        <video className="videoItem" controls controlsList="nodownload" key={item}>
                            <source src={this.videoPath(item)} type="video/mp4" key={item} />
                        </video>
                    </div>
                )}
            </div>
        );
    }
}