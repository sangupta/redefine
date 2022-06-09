import React from 'react';
import { LiveProvider, LiveEditor, LiveError, LivePreview } from 'react-live';

interface Props {
    source: string;
}

const PreviewRenderer = (props: any) => {
    return <div className='preview-render'>{props.children}</div>
}

export default class CodePlayground extends React.Component<Props> {

    render() {
        const {source} = this.props;
        const ComponentLibrary = (window as any).__ComponentLibrary || [];

        if(!ComponentLibrary || ComponentLibrary.length === 0)  {
            // this is read only mode
            return <code>{source}</code>
        }

        return <LiveProvider code={source} scope={{ ...ComponentLibrary }} >
            <LivePreview Component={PreviewRenderer} />
            <LiveError />
            <LiveEditor />
        </LiveProvider>
    }

}
