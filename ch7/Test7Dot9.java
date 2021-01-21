import java.util.ArrayList;

public class Test7Dot9 {

    /**
     * An array structure that can be rotated.
     */
    protected static class CircularArray<T> {
        protected ArrayList<T> arrayList = new ArrayList<T>();
        public boolean add(T e) {
            return arrayList.add(e);
        }
        public void add(int index, T e) {
            arrayList.add(index, e);
        }
        public T get(int index) {
            return arrayList.get(index);
        }
        public void rotate(int n) {
            ArrayList<T> newArrayList = new ArrayList<T>(arrayList.size());
            for (int i = 0; i < arrayList.size(); i++) {
                int rotation = i - n;
                if (rotation < 0) {
                    rotation += arrayList.size();
                }
                newArrayList.add(i, arrayList.get(rotation));
            }
            arrayList = newArrayList;
        }
    }

    public static void main(String[] args) {
        CircularArray<String> array = new CircularArray<>();
        array.add("a");
        array.add("b");
        array.add("c");
        array.add("d");
        array.add("e");
        array.rotate(2);
        System.out.println(array.get(0));
        System.out.println(array.get(1));
        System.out.println(array.get(2));
        System.out.println(array.get(3));
        System.out.println(array.get(4));
    }

}


