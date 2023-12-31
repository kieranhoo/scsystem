import React, { useState, useEffect } from "react";
import {
  View,
  Text,
  StyleSheet,
  Dimensions,
  TouchableOpacity,
  Modal,
  FlatList,
  TouchableWithoutFeedback,
} from "react-native";
import { Colors } from "@/theme/variables";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
import DropDownPicker from "react-native-dropdown-picker";
import { rooms } from "@/services";

interface HeaderProps {}

interface Room {
  room_name: string;
  room_id: string;
}

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const RoomsHeader: React.FC<HeaderProps> = () => {
  const [open, setOpen] = useState(false);
  const [selectedValue, setSelectedValue] = useState<string | null>(null);
  const [items, setItems] = useState<{ label: string; value: string }[]>([]);

  async function fetchData() {
    try {
      const roomData = await rooms.getroominform();
      const roomNames = roomData.data.map((room: Room) => ({
        label: room.room_name,
        value: room.room_id,
      }));
      setItems(roomNames);
    } catch (error) {
      console.error("Error fetching room data:", error);
    }
  }

  useEffect(() => {
    fetchData();
  }, []);

  useEffect(() => {
    if (items.length > 0 && selectedValue === null) {
      setSelectedValue(items[0].value);
    }
  }, [items, selectedValue]);

  return (
    <View style={styles.headerContainer}>
      <DropDownPicker
        open={open}
        items={items}
        value={selectedValue}
        setOpen={setOpen}
        setValue={setSelectedValue}
        setItems={setItems}
        placeholder="Rooms:"
        style={{
          borderRadius: 10,
          borderWidth: 2.5,
          borderColor: "rgba(27, 97, 181, 0.89)",
        }}
        containerStyle={styles.dropdownContainer}
        dropDownContainerStyle={{
          borderTopWidth: 0,
          borderRadius: 10,
          borderWidth: 2.5,
          borderColor: "rgba(27, 97, 181, 0.89)",
        }}
      />

      <TouchableOpacity style={styles.iconContainer}>
        <MaterialCommunityIcons
          name="account-box"
          size={50}
          color={"rgba(27, 97, 181, 0.89)"}
        />
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  headerContainer: {
    flexDirection: "row",
    paddingHorizontal: 5,
    backgroundColor: Colors.WHITE,
    paddingBottom: 0.02 * screenHeight,
    shadowColor: "#000",
    shadowOffset: { width: 1, height: 1 },
    shadowOpacity: 0.4,
    shadowRadius: 3,
    elevation: 5,
    zIndex: 10,
  },
  dropdownContainer: {
    flex: 0.8,
    flexDirection: "row",
    marginLeft: 10,
    marginRight: 5,
    backgroundColor: Colors.WHITE,
  },
  iconContainer: {
    flex: 0.2,
    alignItems: "center",
    justifyContent: "center",
  },
});
