components {
  id: "caution_sprite"
  component: "/main/prefabs/platform/caution_sprite.sprite"
  position {
    x: -0.078816
    y: -0.164486
    z: -3.0
  }
  scale {
    x: 2.0
    y: 2.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"patrol_wall\"\n"
  "mask: \"enemy\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -3.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 14.0\n"
  "  data: 46.5\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
