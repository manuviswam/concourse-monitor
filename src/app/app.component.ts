import { Component } from '@angular/core';
import {PipelineStatus} from "../model/pipeline-status";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  pipelineStatuses: PipelineStatus[];

  constructor() {
    this.pipelineStatuses = [];
    for (let i =0; i<10 ; i++) {
      this.pipelineStatuses.push(new PipelineStatus("id" +i, "Pipeline "+i, i%3 == 0, i));
    }
  }
}
