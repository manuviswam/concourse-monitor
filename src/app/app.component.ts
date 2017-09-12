import { Component } from '@angular/core';
import {PipelineStatus} from "../model/pipeline-status";
import * as Rx from 'rxjs/Rx';
import {PipelineMessage} from "../model/pipeline-message";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  pipelineStatuses: PipelineStatus[];

  constructor() {
    let messageObserver = this.connect("ws://localhost:8000/ws");
    messageObserver.subscribe((val)=> {
      this.pipelineStatuses = JSON.parse(val.data).PipelineStatuses;
    });
  }

  private socket: Rx.Subject<any>;

  public connect(url): Rx.Subject<any> {
    if(!this.socket) {
      this.socket = this.create(url);
    }
    return this.socket;
  }

  private create(url): Rx.Subject<any> {
    let ws = new WebSocket(url);
    let observable = Rx.Observable.create(
      (obs: Rx.Observer<any>) => {
        ws.onmessage = obs.next.bind(obs);
        ws.onerror = obs.error.bind(obs);
        ws.onclose = obs.complete.bind(obs);
        return ws.close.bind(ws);
      }
    );
    let observer = {
      next: (data: Object) => {
        if (ws.readyState === WebSocket.OPEN) {
          ws.send(JSON.stringify(data));
        }
      },
    };
    return Rx.Subject.create(observer, observable);
  }
}
