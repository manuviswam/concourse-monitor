import {PipelineStatus} from "./pipeline-status";

export class PipelineMessage {
  newBuildBroken: boolean;
  newBuildFixed: boolean;
  pipelineStatuses: PipelineStatus[];
}
