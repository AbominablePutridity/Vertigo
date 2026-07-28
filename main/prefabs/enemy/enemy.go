components {
  id: "enemy_sprite"
  component: "/main/prefabs/enemy/enemy.sprite"
  position {
    x: 0.007142857
    y: 0.025
  }
}
components {
  id: "enemy_controller"
  component: "/main/prefabs/enemy/scripts/enemy_controller.script"
}
embedded_components {
  id: "enem"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.0\n"
  "group: \"enemy\"\n"
  "mask: \"default\"\n"
  "mask: \"knife\"\n"
  "mask: \"patrol_wall\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 13.0\n"
  "  data: 43.5\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
