export class PipelineStatus {
  id: string;
  name: string;
  success: boolean;
  order: number;

  constructor(id: string, name: string, success: boolean, order: number) {
    this.id = id;
    this.name = name;
    this.success = success;
    this.order = order;
  }
}
