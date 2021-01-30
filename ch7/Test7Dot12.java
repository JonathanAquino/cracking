import java.util.LinkedList;
import java.util.ArrayList;
import java.util.Collections;

public class Test7Dot12 {

    protected static class Pair<A, B> {
        protected A a;
        protected B b;
        public Pair(A a, B b) {
            this.a = a;
            this.b = b;
        }
    }

    /**
     * A simple hashtable implementation.
     */
    protected static class Hashtable<K, V> {
        protected ArrayList<LinkedList<Pair<K, V>>> buckets;
        // TODO: Provide a parameterless constructor that starts with
        // a reasonable value and maybe auto-grows.
        public Hashtable(int numBuckets) {
            buckets = new ArrayList<LinkedList<Pair<K, V>>>(Collections.nCopies(numBuckets, null));
        }
        public void put(K key, V value) {
            Pair<K, V> pair = getPair(key);
            if (pair != null) {
                pair.b = value;                
                return;
            }
            int index = getIndex(key);
            LinkedList<Pair<K, V>> bucket = buckets.get(index);
            if (bucket == null) {
                bucket = new LinkedList<Pair<K, V>>();
                buckets.add(index, bucket);
            }
            bucket.add(new Pair<K, V>(key, value));
        }
        public void remove(K key) {
            Pair<K, V> pair = getPair(key);
            if (pair != null) {
                // Tombstone.
                pair.a = null;
                pair.b = null;
            }            
        }
        public V get(K key) {
            Pair<K, V> pair = getPair(key);
            if (pair != null) {
                return pair.b;
            }            
            return null;
        }
        protected Pair<K, V> getPair(K key) {
            LinkedList<Pair<K, V>> bucket = buckets.get(getIndex(key));
            if (bucket == null) {
                return null;
            }
            for (Pair<K, V> pair : bucket) {
                if (pair.a == null) {
                    // Tombstone.
                    continue;
                }
                if (pair.a.equals(key)) {
                    return pair;
                }
            }
            return null;
        }
        protected int getIndex(K key) {
            return key.hashCode() % buckets.size();
        }
    }

    public static void main(String[] args) {
        Hashtable<String, String> hashtable = new Hashtable<>(1000);
        hashtable.put("a", "foo");
        hashtable.put("a", "bar");
        hashtable.put("b", "baz");
        System.out.println(hashtable.get("a"));
        System.out.println(hashtable.get("b"));
        System.out.println(hashtable.get("c"));
        hashtable.remove("b");
        System.out.println(hashtable.get("b"));
    }

}


