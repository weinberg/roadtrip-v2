/*
model feature {
  id         String   @id @default(dbgenerated("gen_random_uuid()")) @db.Uuid
  node_id    String   @db.Uuid
  data       Json?
  glyph      String?  @default(" ")
  created_at DateTime @default(now()) @db.Timestamp(6)
  updated_at DateTime @default(now()) @db.Timestamp(6)
  node       node     @relation(fields: [node_id], references: [id], onDelete: Cascade, onUpdate: NoAction)
}

 */

export enum FeatureTypeEnum {
  ROAD = 'ROAD',
  TOWN = 'TOWN',
}

export interface FeatureData {
  type: FeatureTypeEnum;
}

export interface Feature {
  id: string;
  data: FeatureData;
  glyph: string;
}

const name = (parent) => {
  return parent?.data?.name;
};

export default { name };
