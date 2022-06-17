import ReactMarkdown from 'react-markdown';

export default function ComponentDetails({ component }: { component: ComponentDef }) {
    return <ReactMarkdown>{component.description}</ReactMarkdown>
}
