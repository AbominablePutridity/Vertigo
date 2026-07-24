components {
  id: "flight_knife"
  component: "/main/prefabs/knifes/flight_knife.sprite"
  position {
    x: -0.24647707
    y: 2.7859554
  }
}
components {
  id: "flight_knife1"
  component: "/main/prefabs/knifes/scripts/flight_knife.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 0.5\n"
  "friction: 0.05\n"
  "restitution: 0.7\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 3.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 9.5\n"
  "  data: 2.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
