class CreateEquips < ActiveRecord::Migration
  def change
    create_table :equips, id: false do |t|
      t.string :uuid
      t.string :user_id
      t.integer :level
      t.integer :conf_id
      t.string :evolves
      t.string :equips
      t.integer :exp
    end
  end
end
