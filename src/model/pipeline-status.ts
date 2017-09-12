export class PipelineStatus {
  Id: string;
  Name: string;
  Success: boolean;
  Order: number;

  constructor(id: string, name: string, success: boolean, order: number) {
    this.Id = id;
    this.Name = name;
    this.Success = success;
    this.Order = order;
  }
}
