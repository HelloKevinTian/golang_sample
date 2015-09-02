class CreateUsers < ActiveRecord::Migration
  def change
    create_table :users do |t|
      t.string :uuid
      t.integer :level
      t.integer :exp
      t.string :name
      t.boolean :online
    end
  end
end
