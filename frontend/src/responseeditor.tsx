import ReactMarkdown from "react-markdown";
import { Response } from "./data";

export interface Props {
  response: Response;
  onDeleteClicked: (responseId: string) => void;
}

export function ResponseEditor({response, onDeleteClicked}: Props): JSX.Element {
  return <div className="card char-width-20">
    <h3>Response</h3>
    <div className="controls">
      <button className="microtool danger" onClick={() => onDeleteClicked(response.id)}><i className="fa fa-times"></i></button>
    </div>
    <ReactMarkdown children={response.prompt} />
    <br />
    <h4>Output:</h4>
    <ReactMarkdown children={response.text} />
  </div>;
}
